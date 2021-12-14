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
	g := parseGraph(input)
	return walkGraph(g, "start", "end", nil)
}

func part2(input string) int {
	g := parseGraph(input)
	return walkGraph2(g, "start", "end", nil, false)
}

type void = struct{}

func parseGraph(input string) map[string]map[string]void {
	g := make(map[string]map[string]void)

	lines := strutil.SplitLines(input)
	for _, line := range lines {
		if line == "" {
			continue
		}

		s := strings.Split(line, "-")
		if len(s) != 2 {
			panic(fmt.Sprintf("Expected in format 'a-b' but got '%s'", line))
		}

		if _, ok := g[s[0]]; !ok {
			g[s[0]] = make(map[string]void)
		}
		if _, ok := g[s[1]]; !ok {
			g[s[1]] = make(map[string]void)
		}

		g[s[0]][s[1]] = void{}
		g[s[1]][s[0]] = void{}
	}

	return g
}

func walkGraph(g map[string]map[string]void, start string, end string, visited map[string]void) int {
	if start == end {
		return 1
	}

	visitedCopy := make(map[string]void)
	for k := range visited {
		visitedCopy[k] = void{}
	}

	if strings.ToUpper(start) != start {
		visitedCopy[start] = void{}
	}

	sum := 0
	for next := range g[start] {
		if _, ok := visitedCopy[next]; !ok {
			sum += walkGraph(g, next, end, visitedCopy)
		}
	}
	return sum
}

func walkGraph2(g map[string]map[string]void, start string, end string,
	visited map[string]void, visitedTwice bool) int {

	if start == end {
		return 1
	}

	visitedCopy := make(map[string]void)
	for k := range visited {
		visitedCopy[k] = void{}
	}

	if strings.ToUpper(start) != start {
		visitedCopy[start] = void{}
	}

	sum := 0
	for next := range g[start] {
		_, nextVisited := visitedCopy[next]
		if !nextVisited {
			sum += walkGraph2(g, next, end, visitedCopy, visitedTwice)
		} else if !visitedTwice && next != "start" && next != "end" {
			sum += walkGraph2(g, next, end, visitedCopy, true)
		}
	}
	return sum
}
