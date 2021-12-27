package main

import (
	"adventofcode2021/pkg/strutil"
	_ "embed"
	"fmt"
	"log"
	"math"
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

func (v Vector) ManhattanDist() ManhattanDist {
	return ManhattanDist{
		uint(abs(v.X)),
		uint(abs(v.Y)),
		uint(abs(v.Z)),
	}
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func (v Vector) String() string {
	return fmt.Sprintf("V(%d,%d,%d)", v.X, v.Y, v.Z)
}

type ManhattanDist struct {
	X, Y, Z uint
}

func (d ManhattanDist) EuclidDistance() float64 {
	return math.Sqrt(math.Pow(float64(d.X), 2) +
		math.Pow(float64(d.Y), 2) +
		math.Pow(float64(d.Z), 2))
}

type Line struct {
	P1, P2 Point
}

func (l Line) Distance() ManhattanDist {
	return l.P1.Substract(l.P2).ManhattanDist()
}

func calcDistances(ps []Point) map[ManhattanDist][]Line {
	dist := make(map[ManhattanDist][]Line)
	for i := 0; i < len(ps); i++ {
		for j := i + 1; j < len(ps); j++ {
			l := Line{ps[i], ps[j]}
			d := l.Distance()
			dist[d] = append(dist[d], l)
		}
	}
	return dist
}

func parseScanners(input string) [][]Point {
	var result [][]Point
	scanner := -1
	for _, line := range strutil.SplitLines(input) {
		if strings.HasPrefix(line, "--- scanner") {
			result = append(result, make([]Point, 0))
			scanner++
			continue
		} else if line == "" {
			continue
		}

		if scanner >= 0 {
			result[scanner] = append(result[scanner], Parse(line))
		}
	}
	return result
}

func findOverlappingPoints(scan1 []Point, scan2 []Point) (
	over1 map[Point]struct{},
	over2 map[Point]struct{},
	rotation2 int,
	offset2 Vector,
) {
	for rotID := 0; rotID < 24; rotID++ {
		// rotate scanner 2 points
		rotScan2 := make([]Point, 0, len(scan2))
		for _, p := range scan2 {
			rotScan2 = append(rotScan2, p.Rotate(rotID))
		}

		dist1 := calcDistances(scan1)
		dist2 := calcDistances(rotScan2)

		uniqueDist := make(map[ManhattanDist]int)
		for d := range dist1 {
			uniqueDist[d]++
		}
		for d := range dist2 {
			uniqueDist[d]++
		}

		count := 0
		var matchingDist []ManhattanDist
		for d, c := range uniqueDist {
			if c > 1 {
				count++
				matchingDist = append(matchingDist, d)
			}
		}

		// 12 points should have 66 (12 * 11 / 2) matching lines with the same manhattan distance
		if count >= 66 {
			log.Printf("count: %d rotID: %d matchingDist: %v", count, rotID, matchingDist)

			rotation2 = rotID

			// Determine the 12 matching points
			over1 = make(map[Point]struct{})
			for _, d := range matchingDist {
				for _, l := range dist1[d] {
					over1[l.P1] = struct{}{}
					over1[l.P2] = struct{}{}
				}
			}

			over2 = make(map[Point]struct{})
			for _, d := range matchingDist {
				for _, l := range dist2[d] {
					over2[l.P1.Rotate(-rotID)] = struct{}{}
					over2[l.P2.Rotate(-rotID)] = struct{}{}
				}
			}

			// Determine scanner 2 offset vector
			l1 := dist1[matchingDist[0]][0]
			l2 := dist2[matchingDist[0]][0]
			offset2 = l2.P1.Substract(l1.P1)

			break
		}
	}

	return
}
