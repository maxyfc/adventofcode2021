package main

import (
	"adventofcode2021/pkg/strutil"
	_ "embed"
	"fmt"
)

//go:embed input.txt
var measurements string

func main() {
	fmt.Printf("Part 1: %d\n", part1(measurements))
	fmt.Printf("Part 2: %d\n", part2(measurements))
}

func part1(input string) int {
	values := strutil.MustAtoiSlice(strutil.SplitLines(input))
	return countIncrease(values)
}

func part2(input string) int {
	values := strutil.MustAtoiSlice(strutil.SplitLines(input))

	sums := make([]int, 0, len(values)-2)
	for i := 2; i < len(values); i++ {
		sums = append(sums, values[i]+values[i-1]+values[i-2])
	}

	return countIncrease(sums)
}

func countIncrease(values []int) int {
	inc := 0
	prev := values[0]
	for i := 1; i < len(values); i++ {
		curr := values[i]
		if prev < curr {
			inc++
		}
		prev = curr
	}
	return inc
}
