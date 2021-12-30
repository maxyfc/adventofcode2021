package types

import (
	"fmt"
	"testing"
)

func TestRangeIntersect(t *testing.T) {
	tests := []struct {
		r1, r2   Range
		expected bool
	}{
		// No overlap
		{NewRange(0, 2), NewRange(3, 5), false},
		{NewRange(3, 5), NewRange(0, 2), false},
		// Partial overlap
		{NewRange(0, 2), NewRange(1, 3), true},
		{NewRange(1, 3), NewRange(0, 2), true},
		// One inside another
		{NewRange(0, 3), NewRange(1, 2), true},
		{NewRange(1, 2), NewRange(0, 3), true},
		// Boundary tests
		{NewRange(0, 0), NewRange(0, 0), true},
		{NewRange(0, 1), NewRange(1, 2), true},
		{NewRange(1, 2), NewRange(0, 1), true},
		{NewRange(0, 1), NewRange(0, 0), true},
		{NewRange(0, 0), NewRange(0, 1), true},
		{NewRange(0, 1), NewRange(1, 1), true},
		{NewRange(1, 1), NewRange(0, 1), true},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Test-%03d", i+1), func(t *testing.T) {
			output := test.r1.Intersect(test.r2)
			if test.expected != output {
				t.Errorf("r1: %v r2: %v Expected: %v Got: %v",
					test.r1, test.r2, test.expected, output)
			}
		})
	}
}
