package main

import (
	"adventofcode2021/pkg/strutil"
	_ "embed"
	"fmt"
	"math"
)

//go:embed input.txt
var inputData string

func main() {
	fmt.Printf("Part 1: %d\n", part1(inputData))
	fmt.Printf("Part 2: %d\n", part2(inputData))
}

func part1(input string) int {
	s := sum(strutil.SplitLines(input))
	return s.Magnitude()
}

func part2(input string) int {
	lines := strutil.SplitLines(input)

	max := math.MinInt
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines); j++ {
			if i == j {
				continue
			}

			input := []string{lines[i], lines[j]}
			s := sum(input).Magnitude()
			if max < s {
				max = s
			}
		}
	}

	return max
}

func parse(input string) *Pair {
	var root *Pair
	var stack []*Pair
	pos := -1
	for _, c := range input {
		switch c {
		case '[':
			p := &Pair{}
			if pos > -1 {
				p.Parent = stack[pos]
				if stack[pos].LeftSide == nil {
					stack[pos].LeftSide = p
				} else {
					stack[pos].RightSide = p
				}
			} else {
				root = p
			}
			stack = append(stack, p)
			pos++
		case ']':
			if pos > -1 {
				stack[pos] = nil
				stack = stack[:pos]
				pos--
			}
		case ',':
			continue
		default:
			if '0' <= c && c <= '9' {
				if pos > -1 {
					v := Value(int(c - '0'))
					if stack[pos].LeftSide == nil {
						stack[pos].LeftSide = v
					} else {
						stack[pos].RightSide = v
					}
				}
			} else {
				panic(fmt.Sprintf("Unexpected character: %c", c))
			}
		}
	}
	return root
}

func sum(input []string) *Pair {
	var sum *Pair
	for _, line := range input {
		p := parse(line)
		if sum == nil {
			sum = p
		} else {
			sum = add(sum, p)
		}

		for {
			hasExploded := false
			for sum.Explode() {
				hasExploded = true
			}

			hasSplit := sum.Split()

			if !hasExploded && !hasSplit {
				break
			}
		}
	}
	return sum
}

func add(left, right *Pair) *Pair {
	p := &Pair{
		nil,
		left,
		right,
	}
	left.Parent = p
	right.Parent = p
	return p
}

type ValueOrPair interface {
	IsValue() bool
	Value() int
	Left() ValueOrPair
	Right() ValueOrPair
	Magnitude() int
	String() string
}

type Pair struct {
	Parent    *Pair
	LeftSide  ValueOrPair
	RightSide ValueOrPair
}

func (p *Pair) IsValue() bool      { return false }
func (p *Pair) Value() int         { return 0 }
func (p *Pair) Left() ValueOrPair  { return p.LeftSide }
func (p *Pair) Right() ValueOrPair { return p.RightSide }
func (p *Pair) Magnitude() int {
	return p.LeftSide.Magnitude()*3 + p.RightSide.Magnitude()*2
}
func (p *Pair) String() string {
	return fmt.Sprintf("[%s,%s]", p.LeftSide.String(), p.RightSide.String())
}

func (p *Pair) Depth() int {
	if p.Parent == nil {
		return 1
	} else {
		return 1 + p.Parent.Depth()
	}
}

func (p *Pair) CanExplode() bool {
	return p.LeftSide.IsValue() && p.RightSide.IsValue() && p.Depth() > 4
}

func (p *Pair) Explode() bool {
	if !p.CanExplode() {
		if !p.LeftSide.IsValue() && p.LeftSide.(*Pair).Explode() {
			return true
		}
		if !p.RightSide.IsValue() && p.RightSide.(*Pair).Explode() {
			return true
		}
		return false
	}

	walkUpLeft(p, int(p.LeftSide.(Value)))
	walkUpRight(p, int(p.RightSide.(Value)))

	if p.Parent.IsLeft(p) {
		p.Parent.LeftSide = Value(0)
	} else {
		p.Parent.RightSide = Value(0)
	}
	p.Parent = nil
	return true
}

func (p *Pair) Split() bool {
	if p.LeftSide.IsValue() {
		if p.LeftSide.Value() > 9 {
			p.LeftSide = split(p.LeftSide.Value(), p)
			return true
		}
	} else {
		if p.LeftSide.(*Pair).Split() {
			return true
		}
	}

	if p.RightSide.IsValue() {
		if p.RightSide.Value() > 9 {
			p.RightSide = split(p.RightSide.Value(), p)
			return true
		}
	} else {
		if p.RightSide.(*Pair).Split() {
			return true
		}
	}

	return false
}

func (p *Pair) IsLeft(c *Pair) bool {
	return !p.LeftSide.IsValue() && p.LeftSide.(*Pair) == c
}

func (p *Pair) IsRight(c *Pair) bool {
	return !p.RightSide.IsValue() && p.RightSide.(*Pair) == c
}

func walkUpLeft(p *Pair, v int) {
	if p.Parent == nil {
		return
	} else if p.Parent.IsLeft(p) {
		walkUpLeft(p.Parent, v)
	} else if p.Parent.LeftSide.IsValue() {
		p.Parent.LeftSide = addValue(p.Parent.LeftSide, v)
	} else {
		walkDownRight(p.Parent.LeftSide.(*Pair), v)
	}
}

func walkUpRight(p *Pair, v int) {
	if p.Parent == nil {
		return
	} else if p.Parent.IsRight(p) {
		walkUpRight(p.Parent, v)
	} else if p.Parent.RightSide.IsValue() {
		p.Parent.RightSide = addValue(p.Parent.RightSide, v)
	} else {
		walkDownLeft(p.Parent.RightSide.(*Pair), v)
	}
}

func walkDownLeft(p *Pair, v int) {
	if p.LeftSide.IsValue() {
		p.LeftSide = addValue(p.LeftSide, v)
	} else {
		walkDownLeft(p.LeftSide.(*Pair), v)
	}
}

func walkDownRight(p *Pair, v int) {
	if p.RightSide.IsValue() {
		p.RightSide = addValue(p.RightSide, v)
	} else {
		walkDownRight(p.RightSide.(*Pair), v)
	}
}

func addValue(curr ValueOrPair, v int) Value {
	return Value(int(curr.(Value)) + v)
}

func split(v int, parent *Pair) *Pair {
	return &Pair{
		parent,
		Value(int(math.Floor(float64(v) / 2))),
		Value(int(math.Ceil(float64(v) / 2))),
	}
}

type Value int

func (v Value) IsValue() bool      { return true }
func (v Value) Value() int         { return int(v) }
func (v Value) Left() ValueOrPair  { return nil }
func (v Value) Right() ValueOrPair { return nil }
func (v Value) Magnitude() int     { return int(v) }
func (v Value) String() string     { return fmt.Sprintf("%d", v) }
