package main

import (
	util "adventofcode2021/pkg/testutil"
	"fmt"
	"reflect"
	"testing"
)

func TestParts(t *testing.T) {
	testData := `6,10
0,14
9,10
0,3
10,4
4,11
6,0
6,12
4,1
0,13
10,12
3,4
3,0
8,4
1,10
2,14
8,10
9,0

fold along y=7
fold along x=5`

	util.RunTests(t, testData, []util.TestCase{
		{Desc: "Part 1", PartFunc: part1, Expected: 17},
	})
}

func TestParseInput(t *testing.T) {
	testData := `0,0
3,2

fold along y=1
fold along x=2`

	points, folds := parseInput(testData)

	expectedPoints := map[Point]struct{}{
		{0, 0}: {},
		{3, 2}: {},
	}
	expectedFolds := []Fold{
		{1, 1},
		{0, 2},
	}

	if !reflect.DeepEqual(points, expectedPoints) {
		t.Errorf("Expected: %v Got: %v", expectedPoints, points)
	}

	if !reflect.DeepEqual(folds, expectedFolds) {
		t.Errorf("Expected: %v Got: %v", expectedFolds, folds)
	}
}

func TestFoldPoints(t *testing.T) {
	tests := []struct {
		points   map[Point]struct{}
		foldAxis int
		foldLine int
		expected map[Point]struct{}
	}{
		{
			map[Point]struct{}{
				{0, 0}: {},
				{3, 2}: {},
			},
			1, // y axis
			1,
			map[Point]struct{}{
				{0, 0}: {},
				{3, 0}: {},
			},
		}, {
			map[Point]struct{}{
				{0, 0}: {},
				{3, 0}: {},
			},
			0, // x axis
			2,
			map[Point]struct{}{
				{0, 0}: {},
				{1, 0}: {},
			},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			output := foldPoints(test.points, test.foldAxis, test.foldLine)
			if !reflect.DeepEqual(output, test.expected) {
				t.Errorf("Expected: %v Got: %v", test.expected, output)
			}
		})
	}
}

func TestFold(t *testing.T) {
	tests := []struct {
		value, foldLine, expected int
		onfoldLine                bool
	}{
		{0, 1, 0, false},
		{2, 1, 0, false},
		{1, 1, 0, true},
		{3, 1, -1, false},
		{3, 2, 1, false},
		{4, 2, 0, false},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%d %d", test.value, test.foldLine), func(t *testing.T) {
			output, _ := fold(test.value, test.foldLine)
			if output != test.expected {
				t.Errorf("Expected: %d Got: %d", test.expected, output)
			}
		})
	}
}
