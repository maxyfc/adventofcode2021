package main

import (
	"adventofcode2021/day21/partone"
	_ "embed"
	"fmt"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1(6, 3))
	fmt.Printf("Part 2: %d\n", part2(0, 0))
}

func part1(p1Start, p2Start int) int {
	d := partone.NewDeterministicDice()
	_, finalScore := partone.PlayGame(p1Start, p2Start, d, 1000)
	return finalScore
}

func part2(p1Start, p2Start int) int {
	return 0
}
