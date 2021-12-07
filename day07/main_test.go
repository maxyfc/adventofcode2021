package main

import (
	util "adventofcode2021/pkg/testutil"
	"testing"
)

func TestParts(t *testing.T) {
	testData := `16,1,2,0,4,2,7,1,2,14`

	util.RunTests(t, testData, []util.TestCase{
		{Desc: "Part 1", PartFunc: part1, Expected: 37},
		{Desc: "Part 2", PartFunc: part2, Expected: 168},
	})
}
