package strutil

import (
	"fmt"
	"os"
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

func MustReadFile(f string) string {
	input, err := os.ReadFile(f)
	if err != nil {
		panic(fmt.Sprintf("Error reading '%s': %s", f, err))
	}
	return string(input)
}

func MustAtoi(v string) int {
	i, err := strconv.Atoi(v)
	if err != nil {
		panic(fmt.Sprintf("Error converting '%s' to int.", v))
	}
	return i
}

func ConvertToInts(values []string) []int {
	ints := make([]int, 0, len(values))
	for i := 0; i < len(values); i++ {
		ints = append(ints, MustAtoi(values[i]))
	}
	return ints
}
