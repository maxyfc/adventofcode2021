package point

import (
	"fmt"
	"reflect"
	"testing"
)

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

func TestCalcDistances(t *testing.T) {
	tests := []struct {
		ps       []Point
		expected map[Vector][]Line
	}{
		{
			[]Point{{0, 0, 0}, {0, 1, 0}, {1, 0, 0}},
			map[Vector][]Line{
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
				{1, -1, 0}: {
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
			output := ParseScanners(test.input)
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
			scanners := ParseScanners(test.input)
			if len(scanners) != 2 {
				panic("Expected 2 scanners")
			}

			found, over1, over2, rotation2, offset2 := FindOverlap(scanners[0], scanners[1])

			if !found {
				t.Errorf("found: Expected: %v Got: %v", true, found)
			}
			if !reflect.DeepEqual(test.over1, over1) {
				t.Errorf("over1: Expected: %v Got: %v", test.over1, over1)
			}
			if !reflect.DeepEqual(test.over2, over2) {
				t.Errorf("over2: Expected: %v Got: %v", test.over2, over2)
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

func TestPointCompare(t *testing.T) {
	tests := []struct {
		p1, p2   Point
		expected int
	}{
		{Point{0, 0, 0}, Point{0, 0, 0}, 0},
		{Point{0, 0, 0}, Point{0, 0, 1}, -1},
		{Point{0, 0, 0}, Point{0, 1, 0}, -1},
		{Point{0, 0, 0}, Point{1, 0, 0}, -1},
		{Point{0, 0, 0}, Point{0, 1, 1}, -1},
		{Point{0, 0, 0}, Point{1, 0, 1}, -1},
		{Point{0, 0, 0}, Point{1, 1, 0}, -1},
		{Point{0, 0, 0}, Point{1, 1, 1}, -1},
		{Point{0, 0, 1}, Point{0, 0, 0}, 1},
		{Point{0, 1, 0}, Point{0, 0, 0}, 1},
		{Point{1, 0, 0}, Point{0, 0, 0}, 1},
		{Point{0, 1, 1}, Point{0, 0, 0}, 1},
		{Point{1, 0, 1}, Point{0, 0, 0}, 1},
		{Point{1, 1, 0}, Point{0, 0, 0}, 1},
		{Point{1, 1, 1}, Point{0, 0, 0}, 1},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			output := test.p1.Compare(test.p2)
			if test.expected != output {
				t.Errorf("Expected: %v Got: %v", test.expected, output)
			}
		})
	}
}
