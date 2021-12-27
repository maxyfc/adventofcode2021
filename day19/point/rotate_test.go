package point

import (
	"reflect"
	"testing"
)

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
