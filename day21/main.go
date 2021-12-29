package main

import (
	"adventofcode2021/day21/partone"
	"adventofcode2021/day21/parttwo"
	_ "embed"
	"fmt"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1(6, 3))
	fmt.Printf("Part 2: %d\n", part2(6, 3))
}

func part1(p1Start, p2Start int) int {
	d := partone.NewDeterministicDice()
	_, finalScore := partone.PlayGame(p1Start, p2Start, d, 1000)
	return finalScore
}

func part2(p1Start, p2Start int) int {
	results := parttwo.PlayGame(p1Start, p2Start, 21)
	if results[0] > results[1] {
		return results[0]
	} else {
		return results[1]
	}
}
