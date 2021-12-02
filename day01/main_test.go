package main

import (
	util "adventofcode2021/pkg/testutil"
	"testing"
)

func TestParts(t *testing.T) {
	testData := `199
200
208
210
200
207
240
269
260
263`

	util.RunTests(t, testData, []util.TestCase{
		{Desc: "Part 1", PartFunc: part1, Expected: 7},
		{Desc: "Part 2", PartFunc: part2, Expected: 5},
	})
}
