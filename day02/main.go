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

func part1(input string) int {
	lines := strutil.SplitLines(input)

	horizontal := 0
	depth := 0
	for _, l := range lines {
		instructions := strings.Split(l, " ")
		v := strutil.MustAtoi(instructions[1])
		switch instructions[0] {
		case "forward":
			horizontal += v
		case "down":
			depth += v
		case "up":
			depth -= v
		}
	}

	return horizontal * depth
}

func part2(input string) int {
	lines := strutil.SplitLines(input)

	horizontal := 0
	depth := 0
	aim := 0
	for _, l := range lines {
		instructions := strings.Split(l, " ")
		v := strutil.MustAtoi(instructions[1])
		switch instructions[0] {
		case "forward":
			horizontal += v
			depth += aim * v
		case "down":
			aim += v
		case "up":
			aim -= v
		}
	}

	return horizontal * depth
}
