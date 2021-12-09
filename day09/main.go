package main

import (
	"adventofcode2021/pkg/strutil"
	_ "embed"
	"fmt"
	"sort"
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

	var heightMap [][]int
	for _, line := range lines {
		heightMap = append(heightMap, strutil.MustAtoiSlice(strings.Split(line, "")))
	}

	sum := 0
	for _, p := range findLowestPoints(heightMap) {
		sum += p.val + 1
	}
	return sum
}

func part2(input string) int {
	lines := strutil.SplitLines(input)

	var heightMap [][]int
	for _, line := range lines {
		heightMap = append(heightMap, strutil.MustAtoiSlice(strings.Split(line, "")))
	}

	var basins []int
	for _, p := range findLowestPoints(heightMap) {
		basins = append(basins, getBasinSize(p, heightMap))
	}
	sort.Ints(basins)

	n := len(basins)
	return basins[n-1] * basins[n-2] * basins[n-3]
}

type Point struct {
	x   int
	y   int
	val int
}

func findLowestPoints(heightMap [][]int) []Point {
	var points []Point
	for y, row := range heightMap {
		for x := range row {
			curr := heightMap[y][x]

			// up
			if y > 0 && heightMap[y-1][x] <= curr {
				continue
			}

			// down
			if y < len(heightMap)-1 && heightMap[y+1][x] <= curr {
				continue
			}

			// left
			if x > 0 && heightMap[y][x-1] <= curr {
				continue
			}

			// right
			if x < len(row)-1 && heightMap[y][x+1] <= curr {
				continue
			}

			points = append(points, Point{x, y, curr})
		}
	}
	return points
}

func getBasinSize(start Point, heightMap [][]int) int {
	explorePoints := []Point{start}
	history := make(map[Point]struct{})

	maxY := len(heightMap) - 1
	maxX := len(heightMap[0]) - 1

	size := 0
	for len(explorePoints) > 0 {
		p := explorePoints[0]
		explorePoints = append(explorePoints[:0], explorePoints[1:]...)

		if _, ok := history[p]; ok {
			continue // Already explored
		}
		history[p] = struct{}{}

		size++

		if p.val >= 8 {
			continue
		}

		// up
		if p.y > 0 && heightMap[p.y-1][p.x] == p.val+1 {
			explorePoints = append(explorePoints, Point{p.x, p.y - 1, heightMap[p.y-1][p.x]})
		}

		// down
		if p.y < maxY && heightMap[p.y+1][p.x] == p.val+1 {
			explorePoints = append(explorePoints, Point{p.x, p.y + 1, heightMap[p.y+1][p.x]})
		}

		// left
		if p.x > 0 && heightMap[p.y][p.x-1] == p.val+1 {
			explorePoints = append(explorePoints, Point{p.x - 1, p.y, heightMap[p.y][p.x-1]})
		}

		// right
		if p.x < maxX && heightMap[p.y][p.x+1] == p.val+1 {
			explorePoints = append(explorePoints, Point{p.x + 1, p.y, heightMap[p.y][p.x+1]})
		}
	}

	return size
}
