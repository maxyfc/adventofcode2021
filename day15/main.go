package main

import (
	"adventofcode2021/pkg/strutil"
	_ "embed"
	"fmt"
	"math"
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
	grid := parseGrid(input)
	return grid.findPath()
}

func part2(input string) int {
	return 0
}

func parseGrid(input string) *Grid {
	grid := make([][]int, 0, len(input))
	lines := strutil.SplitLines(input)
	for _, line := range lines {
		grid = append(grid, strutil.MustAtoiSlice(strings.Split(line, "")))
	}
	return &Grid{
		grid,
		len(grid[0]) - 1,
		len(grid) - 1,
		make(map[Line]int),
	}
}

type Point struct {
	X, Y int
}

type Line struct {
	start, end Point
}

type Grid struct {
	g    [][]int
	maxX int
	maxY int
	mem  map[Line]int
}

func (g *Grid) findPath() int {
	return g.findPathInternal(Point{0, 0}, Point{g.maxX, g.maxY}, nil)
}

func (g *Grid) findPathInternal(start, end Point, visited map[Point]struct{}) int {
	risk, ok := g.mem[Line{start, end}]
	if ok {
		return risk
	}

	if start == end {
		return 0
	}

	v := make(map[Point]struct{})
	for k := range visited {
		v[k] = struct{}{}
	}
	v[start] = struct{}{}

	var risks []int
	for _, next := range g.getNextPoints(start) {
		if _, ok := v[next]; !ok {
			nextRisk := g.findPathInternal(next, end, v)
			if nextRisk < math.MaxInt {
				nextRisk += g.g[next.Y][next.X]
			}
			risks = append(risks, nextRisk)
		}
	}

	if len(risks) == 0 {
		risk = math.MaxInt // trap itself
	} else {
		sort.Ints(risks)
		risk = risks[0]
	}

	g.mem[Line{start, end}] = risk // remember so that we don't need to search again

	return risk
}

func (g *Grid) getNextPoints(p Point) []Point {
	n := make([]Point, 0, 4)

	// top
	if p.Y > 0 {
		n = append(n, Point{p.X, p.Y - 1})
	}
	// bottom
	if p.Y < g.maxY {
		n = append(n, Point{p.X, p.Y + 1})
	}
	// left
	if p.X > 0 {
		n = append(n, Point{p.X - 1, p.Y})
	}
	// right
	if p.X < g.maxX {
		n = append(n, Point{p.X + 1, p.Y})
	}

	return n
}
