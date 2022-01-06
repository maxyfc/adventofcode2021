package main

import (
	"adventofcode2021/day24/validator"
	"fmt"
	"testing"
)

// First attempt to use brute force using an auto translated input.
// This didn't work as the search space is too big.
// It is used for testing instead.
//go:generate go run gen/gen.go -inputFile=input.txt -outputFile=validator/validator.go

func TestParts(t *testing.T) {
	tests := []struct {
		partFunc func(string) int
	}{
		{part1},
		{part2},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i+1), func(t *testing.T) {
			output := test.partFunc(inputData)
			if !validator.Validate(output) {
				t.Errorf("Output not a valid model number: %d", output)
			}
		})
	}
}
