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
	lines := strutil.SplitLines(input)

	count := 0
	for _, line := range lines {
		splits := strings.Split(line, " | ")
		digitOutputs := strings.Fields(splits[1])
		for _, output := range digitOutputs {
			if len(output) == 7 || len(output) == 3 || len(output) == 4 || len(output) == 2 {
				count++
			}
		}
	}
	return count
}

func part2(input string) int {
	lines := strutil.SplitLines(input)

	sum := 0
	for _, line := range lines {
		splits := strings.Split(line, " | ")
		decodeMap := sortDecodeChars(decode(strings.Fields(splits[0])))
		digitSignals := strings.Fields(splits[1])
		for i, signal := range digitSignals {
			sum += decodeMap[sortChars(signal)] * int(math.Pow10(len(digitSignals)-i-1))
		}
	}

	return sum
}

func sortDecodeChars(decode map[string]int) map[string]int {
	results := make(map[string]int)
	for signal, digit := range decode {
		results[sortChars(signal)] = digit
	}
	return results
}

func sortChars(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func decode(signals []string) map[string]int {
	reverse := make(map[int]string)

	// unique signals
	var fiveSegments []string
	var sixSegments []string
	for _, signal := range signals {
		switch len(signal) {
		case 7:
			reverse[8] = signal
		case 4:
			reverse[4] = signal
		case 3:
			reverse[7] = signal
		case 2:
			reverse[1] = signal
		case 5:
			fiveSegments = append(fiveSegments, signal)
		case 6:
			sixSegments = append(sixSegments, signal)
		}
	}

	// 5 segments
	var signal string
	fiveSegments, signal = findDigit3(fiveSegments, reverse)
	reverse[3] = signal

	fiveSegments, signal = findDigit5(fiveSegments, reverse)
	reverse[5] = signal

	reverse[2] = fiveSegments[0]

	// 6 segments
	sixSegments, signal = findDigit0(sixSegments, reverse)
	reverse[0] = signal

	sixSegments, signal = findDigit9(sixSegments, reverse)
	reverse[9] = signal

	reverse[6] = sixSegments[0]

	// reverse the map
	results := make(map[string]int)
	for digit, signal := range reverse {
		results[signal] = digit
	}
	return results
}

func findDigit3(signals []string, reverse map[int]string) ([]string, string) {
	index := 0
	for i, signal := range signals {
		if countChars(signal, reverse[1]) == 2 {
			index = i
			break
		}
	}
	result := signals[index]
	return append(signals[0:index], signals[index+1:]...), result
}

func findDigit5(signals []string, reverse map[int]string) ([]string, string) {
	search := reverse[4]
	for _, c := range reverse[1] {
		search = strings.Replace(search, string(c), "", 1)
	}

	index := 0
	for i, signal := range signals {
		if countChars(signal, search) == 2 {
			index = i
			break
		}
	}
	result := signals[index]
	return append(signals[0:index], signals[index+1:]...), result
}

func findDigit0(signals []string, reverse map[int]string) ([]string, string) {
	index := 0
	for i, signal := range signals {
		if countChars(signal, reverse[5]) == 4 && countChars(signal, reverse[7]) == 3 {
			index = i
			break
		}
	}
	result := signals[index]
	return append(signals[0:index], signals[index+1:]...), result
}

func findDigit9(signals []string, reverse map[int]string) ([]string, string) {
	index := 0
	for i, signal := range signals {
		if countChars(signal, reverse[5]) == 5 && countChars(signal, reverse[7]) == 3 {
			index = i
			break
		}
	}
	result := signals[index]
	return append(signals[0:index], signals[index+1:]...), result
}

func countChars(str string, chars string) int {
	count := 0
	for _, c := range chars {
		if strings.ContainsRune(str, c) {
			count++
		}
	}
	return count
}
