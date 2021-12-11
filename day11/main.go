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
	grid := parseGrid(input)

	flashes := 0
	for i := 0; i < 100; i++ {
		flashes += step(grid)
	}
	return flashes
}

func part2(input string) int {
	grid := parseGrid(input)

	stepNo := 1
	for step(grid) != 100 {
		stepNo++
	}

	return stepNo
}

func parseGrid(input string) [][]int {
	lines := strutil.SplitLines(input)
	var grid [][]int
	for _, line := range lines {
		var row []int
		for _, c := range strings.Split(line, "") {
			row = append(row, strutil.MustAtoi(string(c)))
		}
		grid = append(grid, row)
	}
	return grid
}

func gridString(grid [][]int) string {
	var s strings.Builder
	for _, row := range grid {
		for _, c := range row {
			s.WriteString(fmt.Sprintf("%d", c))
		}
		s.WriteRune('\n')
	}
	return strings.TrimSpace(s.String())
}

type Point struct {
	X, Y int
}

func step(grid [][]int) int {
	var queue []Point

	flashed := make(map[Point]struct{})

	maxX := len(grid[0]) - 1
	maxY := len(grid) - 1

	count := 0
	for y, row := range grid {
		for x := range row {
			grid[y][x]++
			if grid[y][x] > 9 {
				if _, ok := flashed[Point{x, y}]; !ok {
					count++
					flashed[Point{x, y}] = struct{}{}
					queue = flash(queue, x, y, maxX, maxY)
				}
			}
		}
	}

	for len(queue) > 0 {
		p := queue[0]
		queue = append(queue[0:0], queue[1:]...)
		grid[p.Y][p.X]++
		if grid[p.Y][p.X] > 9 {
			if _, ok := flashed[p]; !ok {
				count++
				flashed[Point{p.X, p.Y}] = struct{}{}
				queue = flash(queue, p.X, p.Y, maxX, maxY)
			}
		}
	}

	for y, row := range grid {
		for x := range row {
			if grid[y][x] > 9 {
				grid[y][x] = 0
			}
		}
	}

	return count
}

func flash(queue []Point, x, y, maxX, maxY int) []Point {
	queue = add(queue, Point{x - 1, y - 1}, maxX, maxY)
	queue = add(queue, Point{x, y - 1}, maxX, maxY)
	queue = add(queue, Point{x + 1, y - 1}, maxX, maxY)
	queue = add(queue, Point{x - 1, y}, maxX, maxY)
	queue = add(queue, Point{x + 1, y}, maxX, maxY)
	queue = add(queue, Point{x - 1, y + 1}, maxX, maxY)
	queue = add(queue, Point{x, y + 1}, maxX, maxY)
	queue = add(queue, Point{x + 1, y + 1}, maxX, maxY)
	return queue
}

func add(queue []Point, p Point, maxX, maxY int) []Point {
	if p.X < 0 || p.Y < 0 || p.X > maxX || p.Y > maxY {
		return queue
	}
	return append(queue, p)
}
