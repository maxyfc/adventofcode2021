package main

import (
	"adventofcode2021/pkg/intutil"
	"adventofcode2021/pkg/strutil"
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var inputData string

func main() {
	fmt.Printf("Part 1: %d\n", part1(inputData))
	fmt.Printf("Part 2: %d\n", part2(inputData))
}

func part1(input string) int {
	return calcCost(input, func(crabPos, pos int) int {
		return intutil.Abs(crabPos - pos)
	})
}

func part2(input string) int {
	return calcCost(input, func(crabPos, pos int) int {
		steps := intutil.Abs(crabPos - pos)
		return int(float64(1+steps) * float64(steps) / 2.0)
	})
}

func calcCost(input string, costFunc func(pos, newPos int) int) int {
	crabsPos := strutil.MustAtoiSlice(strings.Split(input, ","))

	maxPos := 0
	for _, pos := range crabsPos {
		maxPos = intutil.Max(maxPos, pos)
	}

	fuelCosts := make([]int, maxPos+1)
	for pos := range fuelCosts {
		cost := 0
		for _, crabPos := range crabsPos {
			currCost := costFunc(crabPos, pos)
			cost += currCost
		}
		fuelCosts[pos] = cost
	}

	minCost := fuelCosts[0]
	for _, cost := range fuelCosts {
		if cost < minCost {
			minCost = cost
		}
	}

	return minCost
}
