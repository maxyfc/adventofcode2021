package main

import (
	"adventofcode2021/pkg/strutil"
	_ "embed"
	"fmt"
	"log"
	"math"
	"strings"
)

//go:embed input.txt
var inputData string

func main() {
	fmt.Printf("Part 1: %d\n", part1(inputData))
	fmt.Printf("Part 2: %d\n", part2(inputData))
}

func part1(input string) uint64 {
	return stepN(input, 10)
}

func part2(input string) uint64 {
	return stepN(input, 40)
}

func stepN(input string, steps int) uint64 {
	template, rules, count := parse(input)

	for i := 0; i < steps; i++ {
		template, count = step(template, rules, count)
	}

	var min uint64 = math.MaxUint64
	var max uint64 = 0
	for _, c := range count {
		if c < min {
			min = c
		}
		if c > max {
			max = c
		}
	}

	return max - min
}

func step(template map[[2]byte]uint64, rules map[[2]byte]byte, count map[byte]uint64) (
	result map[[2]byte]uint64, newCount map[byte]uint64) {

	result = make(map[[2]byte]uint64)
	newCount = make(map[byte]uint64)

	for k, v := range count {
		newCount[k] = v
	}

	for key, c := range template {
		if insert, ok := rules[key]; ok {
			k1 := [2]byte{key[0], insert}
			k2 := [2]byte{insert, key[1]}
			result[k1] += c
			result[k2] += c
			newCount[insert] += c
		} else {
			log.Printf("Skipped key %s", string(key[:]))
		}
	}
	return
}

func parse(input string) (template map[[2]byte]uint64, rules map[[2]byte]byte, count map[byte]uint64) {
	template = make(map[[2]byte]uint64)
	rules = make(map[[2]byte]byte)
	count = make(map[byte]uint64)

	lines := strutil.SplitLines(input)
	for _, line := range lines {
		if line == "" {
			continue
		} else if strings.Contains(line, "->") {
			s := strings.Split(line, " -> ")
			key := [2]byte{s[0][0], s[0][1]}
			rules[key] = s[1][0]
		} else {
			for i := 0; i < len(line)-1; i++ {
				key := [2]byte{line[i], line[i+1]}
				template[key]++
				if i == 0 {
					count[line[i]]++
				}
				count[line[i+1]]++
			}
		}
	}

	return
}
