package main

import (
	util "adventofcode2021/pkg/testutil"
	"testing"
)

func TestParts(t *testing.T) {
	testData := `5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526`

	util.RunTests(t, testData, []util.TestCase{
		{Desc: "Part 1", PartFunc: part1, Expected: 1656},
		{Desc: "Part 2", PartFunc: part2, Expected: 195},
	})
}

func TestStep(t *testing.T) {
	tests := []struct {
		input    string
		flash    int
		expected string
	}{
		{`5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526`, 0,
			`6594254334
3856965822
6375667284
7252447257
7468496589
5278635756
3287952832
7993992245
5957959665
6394862637`},
		{`6594254334
3856965822
6375667284
7252447257
7468496589
5278635756
3287952832
7993992245
5957959665
6394862637`, 35,
			`8807476555
5089087054
8597889608
8485769600
8700908800
6600088989
6800005943
0000007456
9000000876
8700006848`},
	}
	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			grid := parseGrid(test.input)
			flash := step(grid)
			output := gridString(grid)
			if output != test.expected {
				t.Errorf("Expected:\n%s\nGot:\n%s\n", test.expected, output)
			}
			if flash != test.flash {
				t.Errorf("Expected flash: %d Got: %d", test.flash, flash)
			}
		})
	}
}
