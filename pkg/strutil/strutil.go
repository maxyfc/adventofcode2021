package strutil

import (
	"fmt"
	"strconv"
	"strings"
)

func SplitLines(s string) []string {
	lines := strings.Split(s, "\n")
	for i := 0; i < len(lines); i++ {
		lines[i] = strings.TrimSpace(lines[i])
	}
	return lines
}

func MustAtoi(v string) int {
	i, err := strconv.Atoi(v)
	if err != nil {
		panic(fmt.Sprintf("Error converting '%s' to int.", v))
	}
	return i
}

func MustAtoiSlice(values []string) []int {
	ints := make([]int, 0, len(values))
	for i := 0; i < len(values); i++ {
		ints = append(ints, MustAtoi(values[i]))
	}
	return ints
}
