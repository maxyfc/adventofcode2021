package main

import (
	"adventofcode2021/day23/types"
	_ "embed"
	"fmt"
)

//go:embed input.txt
var inputData string

func main() {
	fmt.Printf("Part 1: %d\n", part1(inputData))
	fmt.Printf("Part 2: %d\n", part2(inputData))
}

func part1(input string) int {
	w := types.NewWorld(2,
		types.PodTypeB,
		types.PodTypeA,
		types.PodTypeC,
		types.PodTypeD,
		types.PodTypeB,
		types.PodTypeC,
		types.PodTypeD,
		types.PodTypeA,
	)

	fmt.Println(w)

	return 0
}

func part2(input string) int {
	return 0
}
