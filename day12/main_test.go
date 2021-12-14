package main

import (
	util "adventofcode2021/pkg/testutil"
	"fmt"
	"reflect"
	"testing"
)

func TestParts(t *testing.T) {
	testData := `start-A
start-b
A-c
A-b
b-d
A-end
b-end`

	util.RunTests(t, testData, []util.TestCase{
		{Desc: "Part 1", PartFunc: part1, Expected: 10},
		{Desc: "Part 2", PartFunc: part2, Expected: 36},
	})
}

func TestParseGraph(t *testing.T) {
	tests := []struct {
		input    string
		expected map[string]map[string]void
	}{
		{
			``, map[string]map[string]void{},
		},
		{
			`start-end`,
			map[string]map[string]void{
				"start": {"end": void{}},
				"end":   {"start": void{}},
			},
		},
		{
			`start-a
start-b
a-c
a-end
b-end`,
			map[string]map[string]void{
				"start": {"a": void{}, "b": void{}},
				"end":   {"a": void{}, "b": void{}},
				"a":     {"start": void{}, "end": void{}, "c": void{}},
				"b":     {"start": void{}, "end": void{}},
				"c":     {"a": void{}},
			},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {
			output := parseGraph(test.input)
			if !reflect.DeepEqual(output, test.expected) {
				t.Errorf("Expected: %v Got: %v", test.expected, output)
			}
		})
	}
}

func TestWalkGraph(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{`start-end`, 1},
		{
			`start-a
a-end`,
			1,
		},
		{
			`start-a
start-b
a-end
b-end`,
			2,
		},
		{
			`start-a
start-b
a-b
a-end
b-end`,
			4,
		},
		{
			`start-a
start-B
a-B
a-end
B-end`,
			5,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {
			g := parseGraph(test.input)
			output := walkGraph(g, "start", "end", nil)
			if output != test.expected {
				t.Errorf("Expected: %v Got: %v", test.expected, output)
			}
		})
	}
}
