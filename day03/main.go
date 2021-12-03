package main

import (
	"adventofcode2021/pkg/strutil"
	"bufio"
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

//go:embed input.txt
var inputData string

func main() {
	fmt.Printf("Part 1: %d\n", part1(inputData))
	fmt.Printf("Part 2: %d\n", part2(inputData))
}

func part1(input string) int {
	bitCount, lineCount := countBits(input)

	gamma := 0
	epsilon := 0
	for p, b := range bitCount {
		v := int(math.Pow(2, float64(len(bitCount)-p-1)))
		if b > lineCount/2 {
			gamma += v
		} else {
			epsilon += v
		}
	}

	return gamma * epsilon
}

func part2(input string) int {
	bits := strutil.SplitLines(input)
	o2 := filter(bits, '1')
	co2 := filter(bits, '0')
	return parseBinary(o2[0]) * parseBinary(co2[0])
}

func countBits(input string) (bitCount []int, lineCount int) {
	s := bufio.NewScanner(strings.NewReader(input))

	bitCount = make([]int, 0, 12)
	lineCount = 0
	for s.Scan() {
		lineCount++
		b := s.Bytes()

		if len(bitCount) < len(b) {
			bitCount = bitCount[:len(b)]
		}

		for i := 0; i < len(b); i++ {
			if b[i] == '1' {
				bitCount[i]++
			}
		}
	}

	return
}

func filter(bits []string, c byte) []string {
	n := len(bits[0])
	for i := 0; i < n && len(bits) > 1; i++ {
		zeros, ones := split(bits, i)
		if c == '0' {
			if len(zeros) <= len(ones) {
				bits = zeros
			} else {
				bits = ones
			}
		} else {
			if len(ones) >= len(zeros) {
				bits = ones
			} else {
				bits = zeros
			}
		}
	}
	return bits
}

func split(bits []string, pos int) (zeros []string, ones []string) {
	for _, b := range bits {
		if b[pos] == '0' {
			zeros = append(zeros, b)
		} else {
			ones = append(ones, b)
		}
	}
	return
}

func parseBinary(s string) int {
	i, err := strconv.ParseInt(s, 2, 32)
	if err != nil {
		panic(fmt.Sprintf("Unable to convert binary string '%s' to int.", s))
	}
	return int(i)
}
