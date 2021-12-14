package main

import (
	"fmt"
	"testing"
)

func TestParts(t *testing.T) {
	testData := ``

	tests := []struct {
		partFunc func(string) int
		expected int
	}{
		{part1, 0},
		{part2, 0},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i+1), func(t *testing.T) {
			output := test.partFunc(testData)
			if output != test.expected {
				t.Errorf("Expected output: %d Got: %d", test.expected, output)
			}
		})
	}
}
