package main

import (
	"fmt"
	"testing"
)

func TestParts(t *testing.T) {
	tests := []struct {
		partFunc func(int, int) int
		expected int
	}{
		{part1, 739785},
		{part2, 444356092776315},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i+1), func(t *testing.T) {
			output := test.partFunc(4, 8)
			if output != test.expected {
				t.Errorf("Expected output: %d Got: %d", test.expected, output)
			}
		})
	}
}
