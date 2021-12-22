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

func TestParse(t *testing.T) {
	tests := []string{
		"[1,2]",
		"[[1,2],3]",
		"[9,[8,7]]",
		"[[1,9],[8,5]]",
		"[[[[1,2],[3,4]],[[5,6],[7,8]]],9]",
		"[[[9,[3,8]],[[0,9],6]],[[[3,7],[4,9]],3]]",
		"[[[[1,3],[5,3]],[[1,3],[8,7]]],[[[4,9],[6,9]],[[8,2],[7,3]]]]",
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			output := parse(test)
			if output.String() != test {
				t.Errorf("Expected: %s Got: %s", test, output.String())
			}
		})
	}
}

func TestParseParent(t *testing.T) {
	output := parse("[[1,2],3]")

	if output.Parent != nil {
		t.Errorf("Root node should have no parent. Got: %v", output.Parent)
	}

	childPair, ok := output.LeftSide.(*Pair)
	if !ok {
		t.Fatalf("Left child node should be a pair. Got: %v", output.LeftSide)
	}
	if childPair.Parent != output {
		t.Errorf("Child node should have a parent. Got: %v", childPair.Parent)
	}
}

func TestPairDepth(t *testing.T) {
	output := parse("[[1,2],3]")

	if output.Depth() != 1 {
		t.Errorf("Root node should have a depth of 1. Got: %v", output.Depth())
	}

	childPair, ok := output.LeftSide.(*Pair)
	if !ok {
		t.Fatalf("Left child node should be a pair. Got: %v", output.LeftSide)
	}
	if childPair.Depth() != 2 {
		t.Errorf("Child node should have a depth of 2. Got: %v", childPair.Depth())
	}
}

func TestPairExplode(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"[[[[[9,8],1],2],3],4]", "[[[[0,9],2],3],4]"},
		{"[7,[6,[5,[4,[3,2]]]]]", "[7,[6,[5,[7,0]]]]"},
		{"[[6,[5,[4,[3,2]]]],1]", "[[6,[5,[7,0]]],3]"},
		{"[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]", "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]"},
		{"[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]", "[[3,[2,[8,0]]],[9,[5,[7,0]]]]"},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			output := parse(test.input)
			if !output.Explode() {
				t.Error("Explode should return true to indicate explosion.")
			}
			if output.String() != test.expected {
				t.Errorf("Expected: %s Got: %s", test.expected, output.String())
			}
		})
	}
}

func TestSplitCalc(t *testing.T) {
	tests := []struct {
		input int
		left  Value
		right Value
	}{
		{10, Value(5), Value(5)},
		{11, Value(5), Value(6)},
		{12, Value(6), Value(6)},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			output := split(test.input, nil)
			if output.Parent != nil {
				t.Errorf("Expected parent should be nil. Got: %v", output.Parent)
			}

			if output.LeftSide != test.left {
				t.Errorf("Expected left side: %v Got: %v", test.left, output.LeftSide)
			}

			if output.RightSide != test.right {
				t.Errorf("Expected right side: %v Got: %v", test.right, output.RightSide)
			}
		})
	}
}
