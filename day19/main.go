package main

import (
	"adventofcode2021/pkg/strutil"
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
	scanners := parseScanners(input)

	var uniqueBeacons []Point
	uniqueBeacons = append(uniqueBeacons, scanners[0]...)

	scanners = append(scanners[:], scanners[1:]...)
	for len(scanners) > 0 {
		scanner := scanners[0]
		scanners = append(scanners[:], scanners[1:]...)

		found, _, over2, rotation2, _ := findOverlap(uniqueBeacons, scanner)
		if found {
			for _, p := range scanner {
				if _, ok := over2[p]; !ok {
					uniqueBeacons = append(uniqueBeacons, p.Rotate(rotation2))
				}
			}
		} else {
			scanners = append(scanners, scanner)
		}

		log.Printf("%d %v", len(uniqueBeacons), uniqueBeacons)
	}

	return len(uniqueBeacons)
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

func (p1 Point) Compare(p2 Point) int {
	if p1.X < p2.X {
		return -1
	} else if p1.X > p2.X {
		return 1
	} else if p1.Y < p2.Y {
		return -1
	} else if p1.Y > p2.Y {
		return 1
	} else if p1.Z < p2.Z {
		return -1
	} else if p1.Z > p2.Z {
		return 1
	} else {
		return 0
	}
}

func (p Point) String() string {
	return fmt.Sprintf("P(%d,%d,%d)", p.X, p.Y, p.Z)
}

type Vector struct {
	X, Y, Z int
}

func (v Vector) String() string {
	return fmt.Sprintf("V(%d,%d,%d)", v.X, v.Y, v.Z)
}

type Line struct {
	P1, P2 Point
}

func calcDistances(ps []Point) map[Vector][]Line {
	dist := make(map[Vector][]Line)
	for i := 0; i < len(ps); i++ {
		for j := i + 1; j < len(ps); j++ {
			p1 := ps[i]
			p2 := ps[j]
			if p1.Compare(p2) > 0 {
				p1, p2 = p2, p1
			}
			l := Line{p1, p2}
			d := p2.Substract(p1)
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

func findOverlap(scan1 []Point, scan2 []Point) (
	found bool,
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

		uniqueVec := make(map[Vector]int)
		for d := range dist1 {
			uniqueVec[d]++
		}
		for d := range dist2 {
			uniqueVec[d]++
		}

		count := 0
		var matchingVec []Vector
		for d, c := range uniqueVec {
			if c > 1 {
				count++
				matchingVec = append(matchingVec, d)
			}
		}

		// 12 points should have 66 (12 * 11 / 2) matching lines with the same manhattan distance
		if count >= 66 {
			found = true
			rotation2 = rotID

			// Determine the 12 matching points
			over1 = make(map[Point]struct{})
			for _, d := range matchingVec {
				for _, l := range dist1[d] {
					over1[l.P1] = struct{}{}
					over1[l.P2] = struct{}{}
				}
			}

			over2 = make(map[Point]struct{})
			for _, d := range matchingVec {
				for _, l := range dist2[d] {
					over2[l.P1.Rotate(-rotID)] = struct{}{}
					over2[l.P2.Rotate(-rotID)] = struct{}{}
				}
			}

			// Determine scanner 2 offset vector
			p1 := dist1[matchingVec[0]][0].P1
			p2 := dist2[matchingVec[0]][0].P1
			offset2 = p1.Substract(p2)

			break
		}
	}

	return
}
