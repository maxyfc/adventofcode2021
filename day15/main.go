package main

import (
	"adventofcode2021/pkg/strutil"
	"container/heap"
	_ "embed"
	"fmt"
	"log"
	"strings"
)

//go:embed input.txt
var inputData string

func main() {
	fmt.Printf("Part 1: %d\n", part1(inputData))
	fmt.Printf("Part 2: %d\n", part2(inputData))
}

func part1(input string) int {
	grid := parseGrid(input, 1)
	grid.debug = true
	return grid.FindPath()
}

func part2(input string) int {
	grid := parseGrid(input, 5)
	return grid.FindPath()
}

func parseGrid(input string, n int) *Grid {
	grid := make([][]int, 0, len(input))
	for _, line := range strutil.SplitLines(input) {
		grid = append(grid, strutil.MustAtoiSlice(strings.Split(line, "")))
	}

	ngrid := expandGrid(grid, n)
	maxX := len(ngrid[0]) - 1
	maxY := len(ngrid) - 1
	return &Grid{
		ngrid,
		maxX,
		maxY,
		make([]*Node, (maxY+1)*(maxX+1)),
		false,
	}
}

func expandGrid(g [][]int, n int) [][]int {
	leny := len(g)
	lenx := len(g[0])
	grid := make([][]int, leny*n)
	for y := 0; y < leny*n; y++ {
		grid[y] = make([]int, lenx*n)
		for x := 0; x < lenx*n; x++ {
			n := g[y%leny][x%lenx] + y/leny + x/lenx
			if n > 9 {
				n = (n-1)%9 + 1
			}
			grid[y][x] = n
		}
	}
	return grid
}

type Point struct {
	X, Y int
}

func (p Point) String() string {
	return fmt.Sprintf("(%d,%d)", p.X, p.Y)
}

type Grid struct {
	g     [][]int
	maxX  int
	maxY  int
	nodes []*Node
	debug bool
}

func (g *Grid) FindPath() int {
	c := 0
	for i, p := range g.Search() {
		if i == 0 {
			continue
		}
		c += g.g[p.Y][p.X]
	}
	return c
}

func (g *Grid) Search() []Point {
	// A* Search algorithm
	openHeap := &NodeHeap{}
	open := make(map[*Node]struct{})
	closed := make(map[*Node]struct{})

	start := g.GetOrCreateNode(Point{0, 0})
	end := g.GetOrCreateNode(Point{g.maxX, g.maxY})
	heap.Init(openHeap)
	heap.Push(openHeap, start)
	open[start] = struct{}{}

	for {
		if _, isEndClosed := closed[end]; len(*openHeap) == 0 || isEndClosed {
			break
		}

		curr := heap.Pop(openHeap).(*Node)
		delete(open, curr)

		closed[curr] = struct{}{}

		for _, neighbor := range curr.GetNeighbors() {
			if _, isClosed := closed[neighbor]; isClosed {
				continue
			} else if _, isOpen := open[neighbor]; !isOpen {
				neighbor.parent = curr
				heap.Push(openHeap, neighbor)
				open[neighbor] = struct{}{}
			} else if neighbor.G() > curr.G()+neighbor.Cost() {
				// neighbor is in the open list
				neighbor.parent = curr
				heap.Init(openHeap)
			}
		}
	}

	if end.parent != nil {
		if g.debug {
			end.DebugPath(false)
		}
		return end.Path()
	}

	return nil
}

func (g *Grid) GetNextPoints(p Point) []Point {
	n := make([]Point, 0, 4)

	// top
	if p.Y > 0 {
		n = append(n, Point{p.X, p.Y - 1})
	}
	// bottom
	if p.Y < g.maxY {
		n = append(n, Point{p.X, p.Y + 1})
	}
	// left
	if p.X > 0 {
		n = append(n, Point{p.X - 1, p.Y})
	}
	// right
	if p.X < g.maxX {
		n = append(n, Point{p.X + 1, p.Y})
	}

	return n
}

func (g *Grid) GetOrCreateNode(p Point) *Node {
	i := p.Y*(g.maxX+1) + p.X
	n := g.nodes[i]
	if n == nil {
		n = &Node{p, nil, g, 0, nil}
		g.nodes[i] = n
	}
	return n
}

func (g *Grid) DebugPath(n *Node, showG bool) {
	points := make(map[Point]struct{})
	var last Point
	for _, p := range n.Path() {
		points[p] = struct{}{}
		last = p
	}

	var s strings.Builder
	for y := 0; y <= g.maxY; y++ {
		for x := 0; x <= g.maxX; x++ {
			i := y*(g.maxX+1) + x
			if showG {
				if _, ok := points[Point{x, y}]; ok {
					if last.X == x && last.Y == y {
						s.WriteString(fmt.Sprintf("*%3d|", g.nodes[i].G()))
					} else {
						s.WriteString(fmt.Sprintf("+%3d|", g.nodes[i].G()))
					}
				} else {
					if g.nodes[i] != nil {
						s.WriteString(fmt.Sprintf("-%3d|", g.nodes[i].G()))
					} else {
						s.WriteString(fmt.Sprintf(" %3d|", g.g[y][x]))
					}
				}
			} else {
				if _, ok := points[Point{x, y}]; ok {
					s.WriteByte('.')
				} else {
					s.WriteString(fmt.Sprintf("%d", g.g[y][x]))
				}
			}
		}

		s.WriteString("\n")

		if showG {
			for x := 0; x <= g.maxX; x++ {
				s.WriteString("-----")
			}
			s.WriteString("\n")
		}
	}

	log.Printf("%v\n%s", n.Path(), s.String())
}

type Node struct {
	Point
	parent *Node
	grid   *Grid

	// Cached g calculation. Only recalculate if the parent is different
	cachedG       int
	cachedGParent *Node
}

// This is the cost function. It is actual cost (g) plus the estimated cost (h)
func (n *Node) F() int {
	return n.G() + n.H()
}

// This is the actual cost from the source to the current point
func (n *Node) G() int {
	if n.parent != n.cachedGParent {
		n.cachedG = n.Cost() + n.parent.G()
		n.cachedGParent = n.parent
	}
	return n.cachedG
}

// This is the estimate cost from the current point to the destination
func (n *Node) H() int {
	return n.grid.maxX - n.Point.X + n.grid.maxY - n.Point.Y // Using manhattan distance for estimation
}

func (n *Node) Cost() int {
	return n.grid.g[n.Point.Y][n.Point.X]
}

func (n *Node) IsDestination() bool {
	return n.Point.X == n.grid.maxX && n.Point.Y == n.grid.maxY
}

func (n *Node) GetNeighbors() []*Node {
	var neighbors []*Node
	for _, p := range n.grid.GetNextPoints(n.Point) {
		n := n.grid.GetOrCreateNode(p)
		neighbors = append(neighbors, n)
	}
	return neighbors
}

func (n *Node) Path() []Point {
	var p []Point
	for curr := n; curr != nil; curr = curr.parent {
		p = append(p, curr.Point)
	}
	for i := 0; i < len(p)/2; i++ {
		p[i], p[len(p)-i-1] = p[len(p)-i-1], p[i]
	}
	return p
}

func (n *Node) DebugPath(showG bool) {
	n.grid.DebugPath(n, showG)
}

// This is min-heap of the node with lowest cost (f)
type NodeHeap []*Node

func (h NodeHeap) Len() int           { return len(h) }
func (h NodeHeap) Less(i, j int) bool { return h[i].F() < h[j].F() }
func (h NodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *NodeHeap) Pop() interface{} {
	old := *h
	l := len(old)
	n := old[l-1]
	*h = old[:l-1]
	return n
}

func (h *NodeHeap) Push(n interface{}) {
	*h = append(*h, n.(*Node))
}
