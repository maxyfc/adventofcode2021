package main

import (
	"adventofcode2021/day22/types"
	"adventofcode2021/pkg/strutil"
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
	boundary := types.NewCubiod(-50, 50, -50, 50, -50, 50)
	return run(input, boundary)
}

func part2(input string) int {
	return run(input, nil)
}

func run(input string, boundary *types.Cubiod) int {
	lines := strutil.SplitLines(input)

	var allCubiods []*types.Cubiod
	for _, line := range lines {
		on, newCubiod := types.ParseLine(line)
		if boundary != nil && !newCubiod.Intersect(boundary) {
			continue
		}

		var cubiods []*types.Cubiod
		for _, c := range allCubiods {
			splits, ok := c.Split(newCubiod)
			if !ok {
				cubiods = append(cubiods, splits...)
				continue
			}
			for _, s := range splits {
				if !newCubiod.Intersect(s) {
					cubiods = append(cubiods, s)
				}
			}
		}

		if on {
			cubiods = append(cubiods, newCubiod)
		}

		allCubiods = cubiods
	}

	count := 0
	for _, c := range allCubiods {
		count += c.Volume()
	}

	return count
}
