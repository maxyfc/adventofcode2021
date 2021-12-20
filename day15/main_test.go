package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParts(t *testing.T) {
	testData := `1163751742
1381373672
2136511328
3694931569
7463417111
1319128137
1359912421
3125421639
1293138521
2311944581`

	tests := []struct {
		partFunc func(string) int
		expected int
	}{
		{part1, 40},
		{part2, 315},
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

func TestParseGrid(t *testing.T) {
	testData := `123
456
789`

	output := parseGrid(testData, 1)
	expected := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	if !reflect.DeepEqual(output.g, expected) {
		t.Errorf("Expected: %v Got: %v", expected, output)
	}
}

func TestFindPath(t *testing.T) {
	tests := []struct {
		grid     string
		expected int
	}{
		{
			`12
11`,
			2,
		},
		{
			`111
122
111`,
			4,
		},
		{
			`1121
1211
1111
1121`,
			6,
		},
		{
			`1163751
1381373
2136511`,
			20,
		},
		{
			`11637517
13813736
21365113
36949315`,
			26,
		},
		{
			`116375174
138137367
213651132
369493156
746341711
131912813`,
			31,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			grid := parseGrid(test.grid, 1)
			result := grid.FindPath()
			if result != test.expected {
				t.Errorf("Expected: %d Got: %d", test.expected, result)
			}
		})
	}
}
