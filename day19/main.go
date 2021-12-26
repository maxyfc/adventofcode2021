package main

import (
	"adventofcode2021/pkg/strutil"
	_ "embed"
	"fmt"
	"math"
	"sort"
	"strings"
)

//go:embed input.txt
var inputData string

func main() {
	fmt.Printf("Part 1: %d\n", part1(inputData))
	fmt.Printf("Part 2: %d\n", part2(inputData))
}

func part1(input string) int {
	return 0
}

func part2(input string) int {
	return 0
}

//go:generate go run gen/main.go -pkgName=main -typeName=Point -output=rotate.go

type Point struct {
	X, Y, Z int
}

func Parse(input string) Point {
	s := strings.Split(input, ",")
	if len(s) != 3 {
		panic(fmt.Sprintf("Should 3 integers separated by comma. Got: %s", input))
	}
	v := strutil.MustAtoiSlice(s)
	return Point{v[0], v[1], v[2]}
}

func (p1 Point) Substract(p2 Point) Vector {
	return Vector{p1.X - p2.X, p1.Y - p2.Y, p1.Z - p2.Z}
}

func (p Point) String() string {
	return fmt.Sprintf("P(%d,%d,%d)", p.X, p.Y, p.Z)
}

type Vector struct {
	X, Y, Z int
}

func (v Vector) Distance() float64 {
	return math.Sqrt(math.Pow(float64(v.X), 2) +
		math.Pow(float64(v.Y), 2) +
		math.Pow(float64(v.Z), 2))
}

func (v Vector) String() string {
	return fmt.Sprintf("V(%d,%d,%d)", v.X, v.Y, v.Z)
}

func calcDistances(ps []Point) []float64 {
	var dist []float64
	for i := 0; i < len(ps); i++ {
		for j := i + 1; j < len(ps); j++ {
			dist = append(dist, ps[i].Substract(ps[j]).Distance())
		}
	}
	sort.Float64s(dist)
	return dist
}
