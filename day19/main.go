package main

import (
	"adventofcode2021/day19/point"
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
	points, _ := matchPoints(input)
	return len(points)
}

func part2(input string) int {
	_, offsets := matchPoints(input)

	offsets = append(offsets, point.Vector{X: 0, Y: 0, Z: 0})
	max := 0
	for i := 0; i < len(offsets); i++ {
		for j := i + 1; j < len(offsets); j++ {
			curr := abs(offsets[i].X-offsets[j].X) +
				abs(offsets[i].Y-offsets[j].Y) +
				abs(offsets[i].Z-offsets[j].Z)
			if max < curr {
				max = curr
			}
		}
	}

	return max
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func matchPoints(input string) (map[point.Point]struct{}, []point.Vector) {
	scanners := point.ParseScanners(input)

	todo := [][]point.Point{scanners[0]}
	scanners = append(scanners[:0], scanners[1:]...)
	done := [][]point.Point{}
	offsets := []point.Vector{}
	for len(todo) > 0 {
		scan1 := todo[0]
		todo = append(todo[:0], todo[1:]...)

		toRemove := make(map[int]struct{})
		for i := 0; i < len(scanners); i++ {
			found, _, _, rotation2, offset2 := point.FindOverlap(scan1, scanners[i])
			if found {
				var scanTodo []point.Point
				for _, p := range scanners[i] {
					scanTodo = append(scanTodo, p.Rotate(rotation2).Offset(offset2))
				}
				todo = append(todo, scanTodo)
				toRemove[i] = struct{}{}
				offsets = append(offsets, offset2)
			}
		}

		scannersCopy := append(scanners[:0:0], scanners...)
		scanners = scanners[0:0]
		for i, s := range scannersCopy {
			if _, exists := toRemove[i]; !exists {
				scanners = append(scanners, s)
			}
		}

		done = append(done, scan1)
	}

	uniquePoints := make(map[point.Point]struct{})
	for _, s := range done {
		for _, p := range s {
			uniquePoints[p] = struct{}{}
		}
	}

	return uniquePoints, offsets
}
