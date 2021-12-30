package main

import (
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
	return 0
}

func part2(input string) int {
	return 0
}
