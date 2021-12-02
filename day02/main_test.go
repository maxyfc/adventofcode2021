package main

import (
	util "adventofcode2021/pkg/testutil"
	"testing"
)

func TestParts(t *testing.T) {
	testData := `forward 5
down 5
forward 8
up 3
down 8
forward 2`

	util.RunTests(t, testData, []util.TestCase{
		{Desc: "Part 1", PartFunc: part1, Expected: 150},
		{Desc: "Part 2", PartFunc: part2, Expected: 900},
	})
}
