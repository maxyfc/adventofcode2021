package main

import (
	util "adventofcode2021/pkg/testutil"
	"testing"
)

func TestParts(t *testing.T) {
	testData := ``

	util.RunTests(t, testData, []util.TestCase{
		{Desc: "Part 1", PartFunc: part1, Expected: 0},
		{Desc: "Part 2", PartFunc: part2, Expected: 0},
	})
}
