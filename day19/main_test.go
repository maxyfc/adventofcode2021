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
			output := test.v.ManhattanDist().EuclidDistance()
			if math.Abs(test.expected-output) > 0.000001 {
				t.Errorf("Expected: %f Got: %f", test.expected, output)
			}
		})
	}
}

func TestCalcDistances(t *testing.T) {
	tests := []struct {
		ps       []Point
		expected map[ManhattanDist][]Line
	}{
		{
			[]Point{{0, 0, 0}, {0, 1, 0}, {1, 0, 0}},
			map[ManhattanDist][]Line{
				{0, 1, 0}: {
					{
						Point{0, 0, 0},
						Point{0, 1, 0},
					},
				},
				{1, 0, 0}: {
					{
						Point{0, 0, 0},
						Point{1, 0, 0},
					},
				},
				{1, 1, 0}: {
					{
						Point{0, 1, 0},
						Point{1, 0, 0},
					},
				},
			},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			output := calcDistances(test.ps)
			if !reflect.DeepEqual(test.expected, output) {
				t.Fatalf("Expected: %v Got: %v", test.expected, output)
			}
		})
	}
}

func TestParseScanners(t *testing.T) {
	tests := []struct {
		input    string
		expected [][]Point
	}{
		{
			`--- scanner 0 ---
404,-588,-901`,
			[][]Point{
				{
					{404, -588, -901},
				},
			},
		},
		{
			`--- scanner 0 ---
404,-588,-901
528,-643,409

--- scanner 1 ---
686,422,578
605,423,415`,
			[][]Point{
				{
					{404, -588, -901},
					{528, -643, 409},
				},
				{
					{686, 422, 578},
					{605, 423, 415},
				},
			},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			output := parseScanners(test.input)
			if !reflect.DeepEqual(test.expected, output) {
				t.Fatalf("Expected %v Got: %v", test.expected, output)
			}
		})
	}
}

func TestFindOverlappingPoints(t *testing.T) {
	tests := []struct {
		input     string
		over1     map[Point]struct{}
		over2     map[Point]struct{}
		rotation2 int
		offset2   Vector
	}{
		{
			`--- scanner 0 ---
404,-588,-901
528,-643,409
-838,591,734
390,-675,-793
-537,-823,-458
-485,-357,347
-345,-311,381
-661,-816,-575
-876,649,763
-618,-824,-621
553,345,-567
474,580,667
-447,-329,318
-584,868,-557
544,-627,-890
564,392,-477
455,729,728
-892,524,684
-689,845,-530
423,-701,434
7,-33,-71
630,319,-379
443,580,662
-789,900,-551
459,-707,401

--- scanner 1 ---
686,422,578
605,423,415
515,917,-361
-336,658,858
95,138,22
-476,619,847
-340,-569,-846
567,-361,727
-460,603,-452
669,-402,600
729,430,532
-500,-761,534
-322,571,750
-466,-666,-811
-429,-592,574
-355,545,-477
703,-491,-529
-328,-685,520
413,935,-424
-391,539,-444
586,-435,557
-364,-763,-893
807,-499,-711
755,-354,-619
553,889,-390`,
			map[Point]struct{}{
				{-618, -824, -621}: {},
				{-537, -823, -458}: {},
				{-447, -329, 318}:  {},
				{404, -588, -901}:  {},
				{544, -627, -890}:  {},
				{528, -643, 409}:   {},
				{-661, -816, -575}: {},
				{390, -675, -793}:  {},
				{423, -701, 434}:   {},
				{-345, -311, 381}:  {},
				{459, -707, 401}:   {},
				{-485, -357, 347}:  {},
			},
			map[Point]struct{}{
				{686, 422, 578}:   {},
				{605, 423, 415}:   {},
				{515, 917, -361}:  {},
				{-336, 658, 858}:  {},
				{-476, 619, 847}:  {},
				{-460, 603, -452}: {},
				{729, 430, 532}:   {},
				{-322, 571, 750}:  {},
				{-355, 545, -477}: {},
				{413, 935, -424}:  {},
				{-391, 539, -444}: {},
				{553, 889, -390}:  {},
			},
			20,
			Vector{68, -1246, -43},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			scanners := parseScanners(test.input)
			if len(scanners) != 2 {
				panic("Expected 2 scanners")
			}

			over1, over2, rotation2, offset2 := findOverlappingPoints(scanners[0], scanners[1])

			if !reflect.DeepEqual(test.over1, over1) {
				t.Errorf("over1: Expected: %v Got: %v", test.over1, over1)
			} else {
				t.Logf("over1: %v", over1)
			}
			if !reflect.DeepEqual(test.over2, over2) {
				t.Errorf("over2: Expected: %v Got: %v", test.over2, over2)
			} else {
				t.Logf("over2: %v", over2)
			}
			if test.rotation2 != rotation2 {
				t.Errorf("rotation2: Expected: %v Got: %v", test.rotation2, rotation2)
			}
			if test.offset2 != offset2 {
				t.Errorf("offset2: Expected: %v Got: %v", test.offset2, offset2)
			}
		})
	}
}
