package main

import (
	"adventofcode2021/day23/types"
	_ "embed"
	"fmt"
	"math"
)

func main() {
	w1 := types.NewWorld(2,
		types.PodTypeD,
		types.PodTypeC,
		types.PodTypeC,
		types.PodTypeA,
		types.PodTypeD,
		types.PodTypeA,
		types.PodTypeB,
		types.PodTypeB,
	)

	fmt.Printf("Part 1: %d\n", cost(w1))

	w2 := types.NewWorld(4,
		types.PodTypeD,
		types.PodTypeD,
		types.PodTypeD,
		types.PodTypeC,
		types.PodTypeC,
		types.PodTypeC,
		types.PodTypeB,
		types.PodTypeA,
		types.PodTypeD,
		types.PodTypeB,
		types.PodTypeA,
		types.PodTypeA,
		types.PodTypeB,
		types.PodTypeA,
		types.PodTypeC,
		types.PodTypeB,
	)

	fmt.Printf("Part 2: %d\n", cost(w2))
}

func cost(w *types.World) int {
	cost, _ := findLowestCost(w)
	fmt.Printf("Cache hit: %d\n", cacheHit)
	return cost
}

var cache map[string]int = make(map[string]int)
var cacheHit int

func findLowestCost(w *types.World) (int, bool) {
	key := w.CacheKey()
	cost, exists := cache[key]
	if exists {
		cacheHit++
		return cost, cost >= 0
	}

	if w.IsSolved() {
		return 0, true
	}

	lowestCost := math.MaxInt
	for _, m := range w.NextMoves() {
		w2 := w.Copy()
		w2.Apply(m)
		c, solved := findLowestCost(w2)
		cost := c + m.Cost()
		if solved && cost < lowestCost {
			lowestCost = cost
		}
	}

	if lowestCost < math.MaxInt {
		cache[key] = lowestCost
		return lowestCost, true
	} else {
		cache[key] = -1
		return -1, false
	}
}
