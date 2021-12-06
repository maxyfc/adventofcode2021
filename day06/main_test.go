package main

import (
	util "adventofcode2021/pkg/testutil"
	"fmt"
	"testing"
)

func TestParts(t *testing.T) {
	testData := `3,4,3,1,2`

	util.RunTests(t, testData, []util.TestCase{
		{Desc: "Part 1", PartFunc: func(i string) int { return int(part1(i)) }, Expected: 5934},
	})
}

func TestSimulate(t *testing.T) {
	tests := []struct {
		input    []int
		days     int
		expected [9]uint64
	}{
		{
			[]int{3, 4, 3, 1, 2},
			0,
			[9]uint64{0, 1, 1, 2, 1, 0, 0, 0, 0},
		},
		{
			[]int{3, 4, 3, 1, 2},
			1,
			[9]uint64{1, 1, 2, 1, 0, 0, 0, 0, 0},
		},
		{
			[]int{3, 4, 3, 1, 2},
			2,
			[9]uint64{1, 2, 1, 0, 0, 0, 1, 0, 1},
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Day(s): %d", test.days), func(t *testing.T) {
			output, _ := simulate(test.input, test.days)
			if output != test.expected {
				t.Errorf("Expected: %v Got: %v", test.expected, output)
			}
		})
	}
}
