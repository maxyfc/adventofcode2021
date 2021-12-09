package main

import (
	util "adventofcode2021/pkg/testutil"
	"testing"
)

func TestParts(t *testing.T) {
	testData := `2199943210
3987894921
9856789892
8767896789
9899965678`

	util.RunTests(t, testData, []util.TestCase{
		{Desc: "Part 1", PartFunc: part1, Expected: 15},
		{Desc: "Part 2", PartFunc: part2, Expected: 1134},
	})
}
