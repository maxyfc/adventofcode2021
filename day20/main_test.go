package main

import (
	"fmt"
	"testing"
)

func TestParts(t *testing.T) {
	testData := `..#.#..#####.#.#.#.###.##.....###.##.#..###.####..#####..#....#..#..##..###..######.###...####..#..#####..##..#.#####...##.#.#..#.##..#.#......#.###.######.###.####...#.##.##..#..#..#####.....#.#....###..#.##......#.....#..#..#..##..#...##.######.####.####.#.#...#.......#..#.#.#...####.##.#......#..#...##.#.##..#...##.#.##..###.#......#.#.......#.#.#.####.###.##...#.....####.#..#..#.##.#....##..#.####....##...##..#...#......#.#.......#.......##..####..#...#.#.#...##..#.#..###..#####........#..####......#..#

#..#.
#....
##..#
..#..
..###`

	tests := []struct {
		partFunc func(string) int
		expected int
	}{
		{part1, 35},
		{part2, 3351},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i+1), func(t *testing.T) {
			output := test.partFunc(testData)
			if output != test.expected {
				t.Errorf("Expected output: %d Got: %d", test.expected, output)
			}
		})
	}
}

func TestImageNew(t *testing.T) {
	input := `#..#.
#....
##..#
..#..
..###`
	output := NewImage(input)
	if output.String() != input {
		t.Errorf("Expected:\n%s\nGot:\n%s", input, output.String())
	}
}

func TestImageEnchance(t *testing.T) {
	tests := []struct {
		alg         []byte
		input       string
		infiniteBit int
		expected    string
	}{
		{
			[]byte(`..#.#..#####.#.#.#.###.##.....###.##.#..###.####..#####..#....#..#..##..###..######.###...####..#..#####..##..#.#####...##.#.#..#.##..#.#......#.###.######.###.####...#.##.##..#..#..#####.....#.#....###..#.##......#.....#..#..#..##..#...##.######.####.####.#.#...#.......#..#.#.#...####.##.#......#..#...##.#.##..#...##.#.##..###.#......#.#.......#.#.#.####.###.##...#.....####.#..#..#.##.#....##..#.####....##...##..#...#......#.#.......#.......##..####..#...#.#.#...##..#.#..###..#####........#..####......#..#`),
			`#..#.
#....
##..#
..#..
..###`,
			0,
			`.##.##.
#..#.#.
##.#..#
####..#
.#..##.
..##..#
...#.#.`,
		},
		{
			[]byte(`###############################################################################################################################################################################################################################################################################################################################################################################################################################################################################################################################.`),
			`#`,
			1,
			`...
...
...`,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			image := NewImage(test.input)
			image.Enhance(test.alg, test.infiniteBit)
			if image.String() != test.expected {
				t.Errorf("Expected:\n%s\nGot:\n%s", test.expected, image.String())
			}
		})
	}
}

func TestGetLookup(t *testing.T) {
	tests := []struct {
		height, width       int
		oldPixels           []bool
		oldHeight, oldWidth int
		infiniteBit         int
		expected            []int
	}{
		{
			3, 3,
			[]bool{true},
			1, 1, 0,
			[]int{1, 2, 4, 8, 16, 32, 64, 128, 256},
		},
		{
			4, 4,
			[]bool{true, true, true, true},
			2, 2, 0,
			[]int{
				1, 3, 6, 4,
				9, 27, 54, 36,
				72, 216, 432, 288,
				64, 192, 384, 256,
			},
		},
		{
			3, 3,
			[]bool{true},
			1, 1, 1,
			[]int{511, 511, 511, 511, 511, 511, 511, 511, 511},
		},

		{
			3, 3,
			[]bool{false},
			1, 1, 1,
			[]int{510, 509, 507, 503, 495, 479, 447, 383, 255},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			for row := 0; row < test.height; row++ {
				for col := 0; col < test.width; col++ {
					output := getLookup(
						row, col, test.oldPixels, test.oldHeight, test.oldWidth, test.infiniteBit)
					expected := test.expected[row*test.height+col]
					if expected != output {
						t.Errorf("Row: %d Col: %d Expected: %d Got: %d", row, col, expected, output)
					}
				}
			}
		})
	}
}
