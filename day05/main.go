package main

import (
	"adventofcode2021/pkg/intutil"
	"adventofcode2021/pkg/strutil"
	_ "embed"
	"fmt"
	"regexp"
)

//go:embed input.txt
var inputData string

func main() {
	fmt.Printf("Part 1: %d\n", part1(inputData))
	fmt.Printf("Part 2: %d\n", part2(inputData))
}

func part1(input string) int {
	return findDangerousPoints(input, false)
}

func part2(input string) int {
	return findDangerousPoints(input, true)
}

func findDangerousPoints(input string, includeDiagonal bool) int {
	lines := parseInput(input)

	grid := make(map[Point]int)
	for _, l := range lines {
		if l.IsHorizontal() {
			min := intutil.Min(l.P1.X, l.P2.X)
			max := intutil.Max(l.P1.X, l.P2.X)
			for i := min; i <= max; i++ {
				p := Point{i, l.P1.Y}
				grid[p] += 1
			}
		} else if l.IsVertical() {
			min := intutil.Min(l.P1.Y, l.P2.Y)
			max := intutil.Max(l.P1.Y, l.P2.Y)
			for i := min; i <= max; i++ {
				p := Point{l.P1.X, i}
				grid[p] += 1
			}
		} else if includeDiagonal && l.IsDiagonal() {
			xSign := intutil.Sign(l.P2.X - l.P1.X)
			ySign := intutil.Sign(l.P2.Y - l.P1.Y)
			steps := intutil.Abs(l.P1.X - l.P2.X)
			for i := 0; i <= steps; i++ {
				p := Point{l.P1.X + xSign*i, l.P1.Y + ySign*i}
				grid[p] += 1
			}
		}
	}

	count := 0
	for _, v := range grid {
		if v >= 2 {
			count++
		}
	}

	return count
}

type Line struct {
	P1, P2 Point
}

func (l *Line) IsHorizontal() bool {
	return l.P1.Y == l.P2.Y
}

func (l *Line) IsVertical() bool {
	return l.P1.X == l.P2.X
}

func (l *Line) IsDiagonal() bool {
	return intutil.Abs(l.P1.X-l.P2.X) == intutil.Abs(l.P1.Y-l.P2.Y)
}

type Point struct {
	X, Y int
}

var ventRegexp *regexp.Regexp = regexp.MustCompile(`(\d+),(\d+) -> (\d+),(\d+)`)

func parseInput(input string) []Line {
	lines := strutil.SplitLines(input)

	var parsedLines []Line
	for _, line := range lines {
		parsedLine := ventRegexp.FindStringSubmatch(line)
		if parsedLine == nil {
			panic(fmt.Sprintf("Cannot parse the following line: %s", line))
		}

		parsedLines = append(parsedLines, Line{
			P1: Point{
				X: strutil.MustAtoi(parsedLine[1]),
				Y: strutil.MustAtoi(parsedLine[2]),
			},
			P2: Point{
				X: strutil.MustAtoi(parsedLine[3]),
				Y: strutil.MustAtoi(parsedLine[4]),
			},
		})
	}
	return parsedLines
}
