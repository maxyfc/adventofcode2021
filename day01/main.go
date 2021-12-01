package main

import (
	"adventofcode2021/pkg/strutil"
	"fmt"
)

func main() {
	input := strutil.MustReadFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func part1(input string) int {
	values := strutil.ConvertToInts(strutil.SplitLines(input))
	return countIncrease(values)
}

func part2(input string) int {
	values := strutil.ConvertToInts(strutil.SplitLines(input))

	sums := make([]int, 0, len(values)-2)
	for i := 2; i < len(values); i++ {
		sums = append(sums, values[i]+values[i-1]+values[i-2])
	}

	return countIncrease(sums)
}

func countIncrease(values []int) int {
	increased := 0
	previous := values[0]
	for i := 1; i < len(values); i++ {
		current := values[i]
		if previous < current {
			increased++
		}
		previous = current
	}
	return increased
}
