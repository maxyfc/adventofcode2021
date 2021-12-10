package main

import (
	"adventofcode2021/pkg/strutil"
	_ "embed"
	"fmt"
	"sort"
)

//go:embed input.txt
var inputData string

func main() {
	fmt.Printf("Part 1: %d\n", part1(inputData))
	fmt.Printf("Part 2: %d\n", part2(inputData))
}

func part1(input string) int {
	lines := strutil.SplitLines(input)

	score := 0
	for _, line := range lines {
		isValid, _, gotChar, _ := parseLine(line)
		if !isValid {
			score += invalidScores[gotChar]
		}
	}

	return score
}

func part2(input string) int {
	lines := strutil.SplitLines(input)

	var scores []int
	for _, line := range lines {
		isValid, _, _, complete := parseLine(line)
		if isValid {
			scores = append(scores, scoreComplete(complete))
		}
	}
	sort.Ints(scores)

	return scores[(len(scores) / 2)]
}

var (
	closingCharMapping = map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
		'<': '>',
	}

	invalidScores = map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}

	validScores = map[rune]int{
		')': 1,
		']': 2,
		'}': 3,
		'>': 4,
	}
)

func parseLine(input string) (bool, rune, rune, string) {
	stack := make([]byte, 0, len(input)) // avoid reallocation
	for _, c := range input {
		switch c {
		case '(':
			fallthrough
		case '[':
			fallthrough
		case '<':
			fallthrough
		case '{':
			stack = append(stack, byte(c))
		case ')':
			fallthrough
		case ']':
			fallthrough
		case '>':
			fallthrough
		case '}':
			peek := stack[len(stack)-1]
			expected := closingCharMapping[rune(peek)]
			if expected == c {
				stack = stack[:len(stack)-1]
			} else {
				return false, expected, c, ""
			}
		default:
			return false, 0, c, ""
		}
	}

	// reverse the stack
	complete := make([]byte, 0, len(stack))
	for i := len(stack) - 1; i >= 0; i-- {
		complete = append(complete, byte(closingCharMapping[rune(stack[i])]))
	}

	return true, 0, 0, string(complete)
}

func scoreComplete(complete string) int {
	score := 0
	for _, c := range complete {
		score = score*5 + validScores[c]
	}
	return score
}
