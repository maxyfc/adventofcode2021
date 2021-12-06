package main

import (
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

func part1(input string) uint64 {
	_, sum := simulate(strutil.MustAtoiSlice(strings.Split(input, ",")), 80)
	return sum
}

func part2(input string) uint64 {
	_, sum := simulate(strutil.MustAtoiSlice(strings.Split(input, ",")), 256)
	return sum
}

func simulate(input []int, days int) ([9]uint64, uint64) {
	// the state array stores the of number fish with internal clocks matching the array index
	// e.g. state[0] - gives the number of fish with 0 day for their internal days
	//      state[1] - gives the number of fish with 1 day for their internal days
	var state [9]uint64

	// init state
	for _, i := range input {
		state[i]++
	}

	for d := 0; d < days; d++ {
		reset := state[0] // these will reset back to 6
		// shift everything down
		for i := 1; i < len(state); i++ {
			state[i-1] = state[i]
		}
		state[6] += reset
		state[8] = reset // new fish
	}

	var sum uint64 = 0
	for i := 0; i < len(state); i++ {
		sum += state[i]
	}
	return state, sum
}
