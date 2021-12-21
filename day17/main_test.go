package main

import (
	"fmt"
	"testing"
)

func TestParts(t *testing.T) {
	tests := []struct {
		partFunc func(int, int, int, int) int
		expected int
	}{
		{part1, 45},
		{part2, 112},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i+1), func(t *testing.T) {
			output := test.partFunc(20, 30, -10, -5)
			if output != test.expected {
				t.Errorf("Expected output: %d Got: %d", test.expected, output)
			}
		})
	}
}
