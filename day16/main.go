package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
)

//go:embed input.txt
var inputData string

func main() {
	fmt.Printf("Part 1: %d\n", part1(inputData))
	fmt.Printf("Part 2: %d\n", part2(inputData))
}

func part1(input string) int64 {
	packet, _ := parse(hexToBin(input))
	queue := []Packet{packet}

	var v int64
	for len(queue) > 0 {
		curr := queue[0]
		queue = append(queue[:0:0], queue[1:]...)
		v += curr.Ver()
		queue = append(queue, curr.Sub()...)
	}

	return v
}

func part2(input string) int64 {
	packet, _ := parse(hexToBin(input))
	return calc(packet)
}

func calc(p Packet) int64 {
	switch p.TID() {
	case 0:
		var sum int64 = 0
		for _, sub := range p.Sub() {
			sum += calc(sub)
		}
		return sum
	case 1:
		var product int64 = 1
		for _, sub := range p.Sub() {
			product *= calc(sub)
		}
		return product
	case 2:
		var min int64 = math.MaxInt64
		for _, sub := range p.Sub() {
			m := calc(sub)
			if min > m {
				min = m
			}
		}
		return min
	case 3:
		var max int64 = math.MinInt64
		for _, sub := range p.Sub() {
			m := calc(sub)
			if max < m {
				max = m
			}
		}
		return max
	case 4:
		return p.Val()
	case 5:
		subs := p.Sub()
		if calc(subs[0]) > calc(subs[1]) {
			return 1
		} else {
			return 0
		}
	case 6:
		subs := p.Sub()
		if calc(subs[0]) < calc(subs[1]) {
			return 1
		} else {
			return 0
		}
	case 7:
		subs := p.Sub()
		if calc(subs[0]) == calc(subs[1]) {
			return 1
		} else {
			return 0
		}
	default:
		panic(fmt.Sprintf("Unknown operation type ID %d", p.TID()))
	}
}

func hexToBin(input string) []byte {
	b := make([]byte, 0, len(input)*4)
	for _, c := range input {
		i, err := strconv.ParseInt(string(c), 16, 64)
		if err != nil {
			panic(fmt.Sprintf("Not a valid hexidecimal digit: '%c'", c))
		}
		b = append(b, []byte(fmt.Sprintf("%04b", i))...)
	}
	return b
}

func parse(input []byte) (Packet, []byte) {
	if len(input) < 6 {
		panic("Missing version and type ID header bytes.")
	}

	version, err := strconv.ParseInt(string(input[:3]), 2, 64)
	if err != nil {
		panic(fmt.Sprintf("Error parsing version %s", input[:3]))
	}

	typeID, err := strconv.ParseInt(string(input[3:6]), 2, 64)
	if err != nil {
		panic(fmt.Sprintf("Error parsing type ID %s", input[3:6]))
	}

	input = input[6:]

	switch typeID {
	case PacketTypeLiteral:
		return parseLiteral(version, input)
	default:
		return parseOperator(version, typeID, input)
	}
}

func parseLiteral(version int64, input []byte) (Packet, []byte) {
	var value int64 = 0
	for {
		group := input[:5]
		input = input[5:]

		v, err := strconv.ParseInt(string(group[1:]), 2, 64)
		if err != nil {
			panic(fmt.Sprintf("Error parsing group %s", group))
		}

		value <<= 4
		value += v

		if group[0] == '0' {
			break
		}
	}

	return &LiteralPacket{
		Version: version,
		TypeID:  PacketTypeLiteral,
		Value:   value,
	}, input
}

func parseOperator(version, typeID int64, input []byte) (Packet, []byte) {
	lenTypeID := input[0]
	input = input[1:]

	switch lenTypeID {
	case '0':
		length, err := strconv.ParseInt(string(input[:15]), 2, 64)
		if err != nil {
			panic(fmt.Sprintf("Invalid length %s", input[:15]))
		}
		input = input[15:]

		var subpackets []Packet
		for subinput := input[:length]; len(subinput) > 0; {
			var packet Packet
			packet, subinput = parse(subinput)
			subpackets = append(subpackets, packet)
		}

		return &OperatorPacket{
			Version:    version,
			TypeID:     typeID,
			Subpackets: subpackets,
		}, input[length:]
	case '1':
		length, err := strconv.ParseInt(string(input[:11]), 2, 64)
		if err != nil {
			panic(fmt.Sprintf("Invalid length %s", input[:11]))
		}
		input = input[11:]

		var subpackets []Packet
		for i := 0; i < int(length); i++ {
			var packet Packet
			packet, input = parse(input)
			subpackets = append(subpackets, packet)
		}

		return &OperatorPacket{
			Version:    version,
			TypeID:     typeID,
			Subpackets: subpackets,
		}, input
	default:
		panic(fmt.Sprintf("Unknown length type ID %c", lenTypeID))
	}
}

const (
	PacketTypeLiteral = 4
)

type Packet interface {
	Ver() int64
	TID() int64
	Val() int64
	Sub() []Packet
}

type LiteralPacket struct {
	Version int64
	TypeID  int64
	Value   int64
}

func (p *LiteralPacket) Ver() int64    { return p.Version }
func (p *LiteralPacket) TID() int64    { return p.TypeID }
func (p *LiteralPacket) Val() int64    { return p.Value }
func (p *LiteralPacket) Sub() []Packet { return nil }

type OperatorPacket struct {
	Version    int64
	TypeID     int64
	Subpackets []Packet
}

func (p *OperatorPacket) Ver() int64    { return p.Version }
func (p *OperatorPacket) TID() int64    { return p.TypeID }
func (p *OperatorPacket) Val() int64    { return 0 }
func (p *OperatorPacket) Sub() []Packet { return p.Subpackets }
