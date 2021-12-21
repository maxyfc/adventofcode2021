package main

import (
	_ "embed"
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1(135, 155, -102, -78))
	fmt.Printf("Part 2: %d\n", part2(135, 155, -102, -78))
}

func part1(x1, x2, y1, y2 int) int {
	hits := scan(x1, x2, y1, y2)
	maxY := math.MinInt
	for _, p := range hits {
		if maxY < p.maxY {
			maxY = p.maxY
		}
	}
	return maxY
}

func part2(x1, x2, y1, y2 int) int {
	return len(scan(x1, x2, y1, y2))
}

func scan(x1, x2, y1, y2 int) []*Probe {
	var hits []*Probe
	for dy := min(y1, y2); dy <= int(math.Abs(float64(min(y1, y2)))); dy++ {
		for dx := 0; dx <= x2; dx++ {
			p := NewProbe(dx, dy)
			for !p.OverShot(x1, x2, y1, y2) {
				p.Step()
				if p.InTarget(x1, x2, y1, y2) {
					hits = append(hits, p)
					break
				}
			}
		}
	}
	return hits
}

type Probe struct {
	origdx, origdy, dx, dy, x, y, maxY int
}

func NewProbe(dx, dy int) *Probe {
	return &Probe{
		dx, dy, dx, dy, 0, 0, 0,
	}
}

func (p *Probe) Step() {
	p.x += p.dx
	p.y += p.dy

	if p.dx > 0 {
		p.dx -= 1
	} else if p.dx < 0 {
		p.dx += 1
	}

	p.dy -= 1

	if p.maxY < p.y {
		p.maxY = p.y
	}
}

func (p *Probe) InTarget(x1, x2, y1, y2 int) bool {
	minX := min(x1, x2)
	maxX := max(x1, x2)
	minY := min(y1, y2)
	maxY := max(y1, y2)
	return minX <= p.x && p.x <= maxX && minY <= p.y && p.y <= maxY
}

func (p *Probe) OverShot(x1, x2, y1, y2 int) bool {
	maxX := max(x1, x2)
	minY := min(y1, y2)
	return p.x > maxX || p.y < minY
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}
