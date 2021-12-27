package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParts(t *testing.T) {
	testData := `--- scanner 0 ---
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
553,889,-390

--- scanner 2 ---
649,640,665
682,-795,504
-784,533,-524
-644,584,-595
-588,-843,648
-30,6,44
-674,560,763
500,723,-460
609,671,-379
-555,-800,653
-675,-892,-343
697,-426,-610
578,704,681
493,664,-388
-671,-858,530
-667,343,800
571,-461,-707
-138,-166,112
-889,563,-600
646,-828,498
640,759,510
-630,509,768
-681,-892,-333
673,-379,-804
-742,-814,-386
577,-820,562

--- scanner 3 ---
-589,542,597
605,-692,669
-500,565,-823
-660,373,557
-458,-679,-417
-488,449,543
-626,468,-788
338,-750,-386
528,-832,-391
562,-778,733
-938,-730,414
543,643,-506
-524,371,-870
407,773,750
-104,29,83
378,-903,-323
-778,-728,485
426,699,580
-438,-605,-362
-469,-447,-387
509,732,623
647,635,-688
-868,-804,481
614,-800,639
595,780,-596

--- scanner 4 ---
727,592,562
-293,-554,779
441,611,-461
-714,465,-776
-743,427,-804
-660,-479,-426
832,-632,460
927,-485,-438
408,393,-506
466,436,-512
110,16,151
-258,-428,682
-393,719,612
-211,-452,876
808,-476,-593
-575,615,604
-485,667,467
-680,325,-822
-627,-443,-432
872,-547,-609
833,512,582
807,604,487
839,-516,451
891,-625,532
-652,-548,-490
30,-46,-14`

	tests := []struct {
		partFunc func(string) int
		expected int
	}{
		{part1, 79},
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

			found, over1, over2, rotation2, offset2 := findOverlap(scanners[0], scanners[1])

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

func Test1(t *testing.T) {
	p1 := Point{1, 1, 1}
	p2 := Point{2, 2, 2}

	if p2.Compare(p1) > 0 {
		p1, p2 = p2, p1
	}

	t.Logf("p2 - p1 = %v", p2.Substract(p1))

	t.Log("test")
}
