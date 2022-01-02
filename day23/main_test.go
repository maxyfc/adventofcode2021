package main

import (
	"adventofcode2021/day23/types"
	"fmt"
	"testing"
)

func TestParts(t *testing.T) {
	tests := []struct {
		data     *types.World
		expected int
	}{
		{
			types.NewWorld(2,
				types.PodTypeB,
				types.PodTypeA,
				types.PodTypeC,
				types.PodTypeD,
				types.PodTypeB,
				types.PodTypeC,
				types.PodTypeD,
				types.PodTypeA,
			),
			12521,
		},
		{
			types.NewWorld(4,
				types.PodTypeB,
				types.PodTypeD,
				types.PodTypeD,
				types.PodTypeA,
				types.PodTypeC,
				types.PodTypeC,
				types.PodTypeB,
				types.PodTypeD,
				types.PodTypeB,
				types.PodTypeB,
				types.PodTypeA,
				types.PodTypeC,
				types.PodTypeD,
				types.PodTypeA,
				types.PodTypeC,
				types.PodTypeA,
			),
			44169,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i+1), func(t *testing.T) {
			output := cost(test.data)
			if output != test.expected {
				t.Errorf("Expected output: %d Got: %d", test.expected, output)
			}
		})
	}
}
