package main

import (
	"fmt"
	"math"
	"reflect"
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
	tests := []struct {
		input    string
		expected Point
	}{
		{"1,1,1", Point{1, 1, 1}},
		{"-1,-1,-1", Point{-1, -1, -1}},
		{"1,2,3", Point{1, 2, 3}},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			output := Parse(test.input)
			if output != test.expected {
				t.Errorf("Expected: %v Got: %v", test.expected, output)
			}
		})
	}
}

func TestRotate(t *testing.T) {
	testData := []Point{
		{-1, -1, 1},
		{-2, -2, 2},
		{-3, -3, 3},
		{-2, -3, 1},
		{5, 6, -4},
		{8, 0, 7},
	}

	expectedData := [][]Point{
		{
			{1, -1, 1},
			{2, -2, 2},
			{3, -3, 3},
			{2, -1, 3},
			{-5, 4, -6},
			{-8, -7, 0},
		},
		{
			{-1, -1, -1},
			{-2, -2, -2},
			{-3, -3, -3},
			{-1, -3, -2},
			{4, 6, 5},
			{-7, 0, 8},
		},
		{
			{1, 1, -1},
			{2, 2, -2},
			{3, 3, -3},
			{1, 3, -2},
			{-4, -6, 5},
			{7, 0, 8},
		},
		{
			{1, 1, 1},
			{2, 2, 2},
			{3, 3, 3},
			{3, 1, 2},
			{-6, -4, -5},
			{0, 7, -8},
		},
	}

	result := make(map[int]int) // expectedData index -> rotation id
	for id := 0; id < 24; id++ {
		output := make([]Point, 0, len(testData))
		for _, p := range testData {
			output = append(output, p.Rotate(id))
		}

		found := -1
		for i, expected := range expectedData {
			if reflect.DeepEqual(output, expected) {
				found = i
				break
			}
		}

		if found > -1 {
			if prevId, exists := result[found]; exists {
				t.Fatalf("Duplicate rotation result detected: %#v Got rotation IDs: %d, %d",
					expectedData[found], prevId, id)
			}

			result[found] = id
		}
	}

	if len(result) != len(expectedData) {
		t.Errorf("Missing one or more expected rotations. Current result: %v", result)
	} else {
		t.Logf("Final result: expectedData index -> rotation id: %v", result)
	}
}

func TestSubstract(t *testing.T) {
	tests := []struct {
		p1       Point
		p2       Point
		expected Vector
	}{
		{
			Point{0, 0, 0},
			Point{0, 0, 0},
			Vector{0, 0, 0},
		},
		{
			Point{1, 2, 3},
			Point{0, 0, 0},
			Vector{1, 2, 3},
		},
		{
			Point{0, 0, 0},
			Point{1, 2, 3},
			Vector{-1, -2, -3},
		},
		{
			Point{6, 5, 4},
			Point{1, 2, 3},
			Vector{5, 3, 1},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			output := test.p1.Substract(test.p2)
			if test.expected != output {
				t.Errorf("Expected: %s Got: %s", test.expected, output)
			}
		})
	}
}

func TestVectorDistance(t *testing.T) {
	tests := []struct {
		v        Vector
		expected float64
	}{
		{Vector{0, 0, 0}, 0},
		{Vector{1, 1, 1}, 1.732051},
		{Vector{0, 3, 4}, 5},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			output := test.v.Distance()
			if math.Abs(test.expected-output) > 0.000001 {
				t.Errorf("Expected: %f Got: %f", test.expected, output)
			}
		})
	}
}

func TestCalcDistances(t *testing.T) {
	tests := []struct {
		ps       []Point
		expected []float64
	}{
		{
			[]Point{{0, 0, 0}, {0, 1, 0}, {1, 0, 0}},
			[]float64{1, 1, 1.414214},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			output := calcDistances(test.ps)
			if len(test.expected) != len(output) {
				t.Fatalf("Expected length: %d Got: %d", len(test.expected), len(output))
			}
			for j := range test.expected {
				if math.Abs(test.expected[j]-output[j]) > 0.000001 {
					t.Errorf("Element %d Expected: %f Got: %f", j, test.expected[j], output[j])
				}
			}
		})
	}
}
