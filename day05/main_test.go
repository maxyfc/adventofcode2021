package main

import (
	util "adventofcode2021/pkg/testutil"
	"reflect"
	"testing"
)

func TestParts(t *testing.T) {
	testData := `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`

	util.RunTests(t, testData, []util.TestCase{
		{Desc: "Part 1", PartFunc: part1, Expected: 5},
		{Desc: "Part 2", PartFunc: part2, Expected: 12},
	})
}

func TestParseInput(t *testing.T) {
	testData := `0,9 -> 5,9
8,0 -> 0,8`

	expectedLines := []Line{
		{
			P1: Point{0, 9},
			P2: Point{5, 9},
		},
		{
			P1: Point{8, 0},
			P2: Point{0, 8},
		},
	}

	lines := parseInput(testData)

	if len(lines) != len(expectedLines) {
		t.Fatalf("Expected length: %d Got: %d", len(expectedLines), len(lines))
	}

	if !reflect.DeepEqual(lines, expectedLines) {
		t.Errorf("Expected: %v Got: %v", expectedLines, lines)
	}
}
