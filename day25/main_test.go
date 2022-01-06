package main

import (
	"fmt"
	"testing"
)

func TestParts(t *testing.T) {
	testData := `v...>>.vv>
.vv>>.vv..
>>.>v>...v
>>v>>.>.v.
v>v.vv.v..
>.>>..v...
.vv..>.>v.
v.v..>>v.v
....v..v.>`

	tests := []struct {
		partFunc func(string) int
		expected int
	}{
		{part1, 58},
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

func TestParse(t *testing.T) {
	testData := `v...>>.vv>
.vv>>.vv..
>>.>v>...v
>>v>>.>.v.
v>v.vv.v..
>.>>..v...
.vv..>.>v.
v.v..>>v.v
....v..v.>`

	g := parse(testData)
	if g.String() != testData {
		t.Errorf("Expected:\n%s\nGot:\n%s", testData, g.String())
	}
}

func TestStep(t *testing.T) {
	tests := []struct {
		data     string
		expected string
		moved    bool
	}{
		{
			`..........
.>v....v..
.......>..
..........`,
			`..........
.>........
..v....v>.
..........`,
			true,
		},
		{
			`..>>v>vv..
..v.>>vv..
..>>v>>vv.
..>>>>>vv.
v......>vv
v>v....>>v
vvv.....>>
>vv......>
.>v.vv.v..`,
			`..>>v>vv..
..v.>>vv..
..>>v>>vv.
..>>>>>vv.
v......>vv
v>v....>>v
vvv.....>>
>vv......>
.>v.vv.v..`,
			false,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%03d", i+1), func(t *testing.T) {
			g := parse(test.data)
			moved := g.Step()
			if moved != test.moved {
				t.Errorf("Expected moved: %v Got: %v", test.moved, moved)
			}
			output := g.String()
			if output != test.expected {
				t.Errorf("Expected output: %s Got: %s", test.expected, output)
			}
		})
	}
}
