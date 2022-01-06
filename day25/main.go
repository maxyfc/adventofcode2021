package main

import (
	"adventofcode2021/pkg/strutil"
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var inputData string

func main() {
	fmt.Printf("%d\n", part1(inputData))
	// There is no part 2
}

func part1(input string) int {
	g := parse(input)

	steps := 0
	for {
		steps++
		if !g.Step() {
			break
		}
	}

	return steps
}

func part2(input string) int {
	return 0
}

type grid struct {
	data   []byte
	width  int
	height int
}

func (g *grid) Step() bool {
	moved := false
	for _, move := range []byte{'>', 'v'} {
		new := make([]byte, g.width*g.height)
		for h := 0; h < g.height; h++ {
			for w := 0; w < g.width; w++ {
				switch g.data[h*g.width+w] {
				case '>':
					if move == '>' && g.data[h*g.width+(w+1)%g.width] == '.' {
						new[h*g.width+w] = '.'
						new[h*g.width+(w+1)%g.width] = '>'
						moved = true
					} else {
						new[h*g.width+w] = '>'
					}
				case 'v':
					if move == 'v' && g.data[((h+1)%g.height)*g.width+w] == '.' {
						new[h*g.width+w] = '.'
						new[((h+1)%g.height)*g.width+w] = 'v'
						moved = true
					} else {
						new[h*g.width+w] = 'v'
					}
				case '.':
					if new[h*g.width+w] == 0 {
						new[h*g.width+w] = '.'
					}
				default:
					panic(fmt.Sprintf("Invalid char '%c' at %d, %d", g.data[h*g.width+w], w, h))
				}
			}
		}
		g.data = new
	}

	return moved
}

func (g *grid) String() string {
	var s strings.Builder
	for h := 0; h < g.height; h++ {
		if h > 0 {
			s.WriteByte('\n')
		}
		for w := 0; w < g.width; w++ {
			s.WriteByte(g.data[h*g.width+w])
		}
	}
	return s.String()
}

func parse(input string) *grid {
	lines := strutil.SplitLines(input)
	if len(lines) == 0 {
		panic("Input has no lines")
	}

	g := &grid{
		width:  len(lines[0]),
		height: len(lines),
	}
	g.data = make([]byte, g.width*g.height)

	for h, line := range lines {
		for w, d := range []byte(line) {
			g.data[h*g.width+w] = d
		}
	}

	return g
}
