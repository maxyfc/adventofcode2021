package main

import "testing"

var input string = `199
200
208
210
200
207
240
269
260
263`

func TestParts(t *testing.T) {
	tests := []struct {
		desc     string
		partFunc func(string) int
		expected int
	}{
		{"Part 1", part1, 7},
		{"Part 2", part2, 5},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			output := test.partFunc(input)
			if output != test.expected {
				t.Errorf("Expected output: %d Got: %d", test.expected, output)
			}
		})
	}
}
