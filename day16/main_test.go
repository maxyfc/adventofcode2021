package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParts(t *testing.T) {
	tests := []struct {
		data     string
		partFunc func(string) int64
		expected int64
	}{
		{"8A004A801A8002F478", part1, 16},
		{"620080001611562C8802118E34", part1, 12},
		{"C0015000016115A2E0802F182340", part1, 23},
		{"A0016C880162017C3686B18A3D4780", part1, 31},
		{"C200B40A82", part2, 3},
		{"04005AC33890", part2, 54},
		{"880086C3E88112", part2, 7},
		{"CE00C43D881120", part2, 9},
		{"F600BC2D8F", part2, 0},
		{"9C005AC2F8F0", part2, 0},
		{"9C0141080250320F1802104A08", part2, 1},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i+1), func(t *testing.T) {
			output := test.partFunc(test.data)
			if output != test.expected {
				t.Errorf("Expected output: %d Got: %d", test.expected, output)
			}
		})
	}
}

func TestHexToBin(t *testing.T) {
	tests := []struct {
		input    string
		expected []byte
	}{
		{"", []byte("")},
		{"0", []byte("0000")},
		{"1", []byte("0001")},
		{"2", []byte("0010")},
		{"3", []byte("0011")},
		{"4", []byte("0100")},
		{"5", []byte("0101")},
		{"6", []byte("0110")},
		{"7", []byte("0111")},
		{"8", []byte("1000")},
		{"9", []byte("1001")},
		{"A", []byte("1010")},
		{"B", []byte("1011")},
		{"C", []byte("1100")},
		{"D", []byte("1101")},
		{"E", []byte("1110")},
		{"F", []byte("1111")},
		{"D2FE28", []byte("110100101111111000101000")},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			output := hexToBin(test.input)
			if !reflect.DeepEqual(test.expected, output) {
				t.Errorf("Expected: %s Got: %s", test.expected, output)
			}
		})
	}
}

func TestParseLiteral(t *testing.T) {
	input := []byte("110100101111111000101000")
	expected := &LiteralPacket{
		Version: 6,
		TypeID:  PacketTypeLiteral,
		Value:   2021,
	}

	output, _ := parse(input)
	if !reflect.DeepEqual(expected, output) {
		t.Errorf("Expected: %#v Got: %#v", expected, output)
	}
}

func TestParseOperator(t *testing.T) {
	tests := []struct {
		input    []byte
		expected *OperatorPacket
	}{
		{
			[]byte("00111000000000000110111101000101001010010001001000000000"),
			&OperatorPacket{
				Version: 1,
				TypeID:  6,
				Subpackets: []Packet{
					&LiteralPacket{
						Version: 6,
						TypeID:  PacketTypeLiteral,
						Value:   10,
					},
					&LiteralPacket{
						Version: 2,
						TypeID:  PacketTypeLiteral,
						Value:   20,
					},
				},
			},
		},
		{
			[]byte("11101110000000001101010000001100100000100011000001100000"),
			&OperatorPacket{
				Version: 7,
				TypeID:  3,
				Subpackets: []Packet{
					&LiteralPacket{
						Version: 2,
						TypeID:  PacketTypeLiteral,
						Value:   1,
					},
					&LiteralPacket{
						Version: 4,
						TypeID:  PacketTypeLiteral,
						Value:   2,
					},
					&LiteralPacket{
						Version: 1,
						TypeID:  PacketTypeLiteral,
						Value:   3,
					},
				},
			},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			checkParseOperator(t, test.input, test.expected)
		})
	}
}

func checkParseOperator(t *testing.T, input []byte, expected *OperatorPacket) {
	t.Helper()

	output, _ := parse(input)
	operator, ok := output.(*OperatorPacket)
	if !ok {
		t.Errorf("Expected *OperatorPackget Got: %t", output)
	}

	if expected.Version != operator.Version {
		t.Errorf("Expected Version: %d Got: %d", expected.Version, operator.Version)
	}

	if expected.TypeID != operator.TypeID {
		t.Errorf("Expected TypeID: %d Got: %d", expected.TypeID, operator.TypeID)
	}

	if len(expected.Subpackets) != len(operator.Subpackets) {
		t.Fatalf("Expected len(Subpackets): %d Got: %d", len(expected.Subpackets), len(operator.Subpackets))
	}

	for i, subpackge := range expected.Subpackets {
		if !reflect.DeepEqual(subpackge, operator.Subpackets[i]) {
			t.Errorf("Subpacket %d Expected: %#v Got: %#v", i, subpackge, operator.Subpackets[i])
		}
	}
}
