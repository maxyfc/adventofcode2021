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

	fmt.Printf("Part 2:\n")
	part2(inputData)
}

func part1(input string) int {
	points, folds := parseInput(input)
	points = foldPoints(points, folds[0].Axis, folds[0].Line)
	return len(points)
}

func part2(input string) {
	points, folds := parseInput(input)

	for _, f := range folds {
		points = foldPoints(points, f.Axis, f.Line)
	}

	maxX := 0
	maxY := 0
	for p := range points {
		if p.X > maxX {
			maxX = p.X
		}
		if p.Y > maxY {
			maxY = p.Y
		}
	}

	grid := make([][]byte, maxY+1)
	for y := 0; y <= maxY; y++ {
		row := make([]byte, maxX+1)
		for x := range row {
			row[x] = byte('.')
		}
		grid[y] = row
	}

	for p := range points {
		grid[p.Y][p.X] = '#'
	}

	for _, row := range grid {
		fmt.Println(string(row))
	}
}

type Point struct {
	X, Y int
}

type Fold struct {
	Axis, Line int
}

func parseInput(input string) (map[Point]struct{}, []Fold) {
	lines := strutil.SplitLines(input)

	p := make(map[Point]struct{})
	var f []Fold
	for _, line := range lines {
		if line == "" {
			continue
		} else if strings.HasPrefix(line, "fold along x=") {
			x := strings.TrimPrefix(line, "fold along x=")
			f = append(f, Fold{0, strutil.MustAtoi(x)})
		} else if strings.HasPrefix(line, "fold along y=") {
			y := strings.TrimPrefix(line, "fold along y=")
			f = append(f, Fold{1, strutil.MustAtoi(y)})
		} else {
			splits := strings.Split(line, ",")
			if len(splits) != 2 {
				panic(fmt.Sprintf("Expected in the format x,y. Got: '%s'", line))
			}
			p[Point{strutil.MustAtoi(splits[0]), strutil.MustAtoi(splits[1])}] = struct{}{}
		}
	}

	return p, f
}

func foldPoints(ps map[Point]struct{}, foldAxis int, foldLine int) map[Point]struct{} {
	result := make(map[Point]struct{})

	for p := range ps {
		if foldAxis == 0 {
			// x axis
			x, onFoldLine := fold(p.X, foldLine)
			if !onFoldLine && x >= 0 {
				result[Point{x, p.Y}] = struct{}{}
			}
		} else {
			// y axis
			y, onFoldLine := fold(p.Y, foldLine)
			if !onFoldLine && y >= 0 {
				result[Point{p.X, y}] = struct{}{}
			}
		}
	}

	return result
}

func fold(value int, foldLine int) (int, bool) {
	if value < foldLine {
		return value, false
	} else if value == foldLine {
		return 0, true
	} else {
		return -value + foldLine*2, false
	}
}
