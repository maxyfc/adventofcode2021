package main

import (
	util "adventofcode2021/pkg/testutil"
	"testing"
)

func TestParts(t *testing.T) {
	testData := `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`

	util.RunTests(t, testData, []util.TestCase{
		{Desc: "Part 1", PartFunc: part1, Expected: 198},
		{Desc: "Part 2", PartFunc: part2, Expected: 230},
	})
}
