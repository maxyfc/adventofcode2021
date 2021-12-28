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
	fmt.Printf("Part 1: %d\n", part1(inputData))
	fmt.Printf("Part 2: %d\n", part2(inputData))
}

func part1(input string) int {
	return enhance(input, 2)
}

func part2(input string) int {
	return enhance(input, 50)
}

func enhance(input string, n int) int {
	lines := strutil.SplitLines(input)

	alg := lines[0]
	img := NewImage(strings.Join(lines[2:], "\n"))

	flipInfinitBit := alg[0] == '#'
	infiniteBit := 0
	for i := 0; i < n; i++ {
		if flipInfinitBit {
			infiniteBit = i % 2
		}
		img.Enhance([]byte(alg), infiniteBit)
	}

	return img.Count()
}

type Image struct {
	pixels        []bool
	height, width int
}

func NewImage(input string) *Image {
	lines := strutil.SplitLines(input)
	i := &Image{}
	if len(lines) > 0 {
		i.height = len(lines)
		i.width = len(lines[0])
		i.pixels = parsePixels(i.height, i.width, lines)
	}
	return i
}

func (i *Image) Enhance(alg []byte, infiniteBit int) {
	old := i.pixels
	newHeight := i.height + 2
	newWidth := i.width + 2
	new := make([]bool, (newHeight)*(newWidth))

	for row := 0; row < newHeight; row++ {
		for col := 0; col < newWidth; col++ {
			lookup := getLookup(row, col, old, i.height, i.width, infiniteBit)
			new[row*newHeight+col] = alg[lookup] == '#'
		}
	}

	i.pixels = new
	i.height = newHeight
	i.width = newWidth
}

func (i *Image) Count() int {
	c := 0
	for _, b := range i.pixels {
		if b {
			c++
		}
	}
	return c
}

func (i *Image) String() string {
	var s strings.Builder
	for row := 0; row < i.height; row++ {
		if row > 0 {
			s.WriteByte('\n')
		}
		for col := 0; col < i.width; col++ {
			if i.pixels[row*i.height+col] {
				s.WriteByte('#')
			} else {
				s.WriteByte('.')
			}
		}
	}
	return s.String()
}

func parsePixels(height, width int, lines []string) []bool {
	pixels := make([]bool, height*width)
	for row, line := range lines {
		for col, char := range line {
			switch char {
			case '#':
				pixels[row*height+col] = true
			case '.':
				pixels[row*height+col] = false
			default:
				panic(fmt.Sprintf("Invalid char: %c", char))
			}
		}
	}
	return pixels
}

func getLookup(row, col int, oldPixels []bool, oldHeight, oldWidth int, infiniteBit int) int {
	lookup := 0
	for rowOffset := -1; rowOffset <= 1; rowOffset++ {
		for colOffset := -1; colOffset <= 1; colOffset++ {
			oldRow := row + rowOffset - 1
			oldCol := col + colOffset - 1

			var bit int
			if oldRow < 0 || oldCol < 0 || oldRow >= oldHeight || oldCol >= oldWidth {
				bit = infiniteBit
			} else if oldPixels[oldRow*oldHeight+oldCol] {
				bit = 1
			} else {
				bit = 0
			}

			lookup = lookup << 1
			lookup += bit
		}
	}

	return lookup
}
