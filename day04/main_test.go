package main

import (
	util "adventofcode2021/pkg/testutil"
	"reflect"
	"testing"
)

var testData string = `7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
8  2 23  4 24
21  9 14 16  7
6 10  3 18  5
1 12 20 15 19

3 15  0  2 22
9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
2  0 12  3  7`

var expected = BingoGame{
	DrawnNumbers: []int{
		7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1,
	},
	Boards: []*Board{
		{
			Numbers: [][]int{
				{22, 13, 17, 11, 0},
				{8, 2, 23, 4, 24},
				{21, 9, 14, 16, 7},
				{6, 10, 3, 18, 5},
				{1, 12, 20, 15, 19},
			},
			Rows: []map[int]bool{
				{22: true, 13: true, 17: true, 11: true, 0: true},
				{8: true, 2: true, 23: true, 4: true, 24: true},
				{21: true, 9: true, 14: true, 16: true, 7: true},
				{6: true, 10: true, 3: true, 18: true, 5: true},
				{1: true, 12: true, 20: true, 15: true, 19: true},
			},
			Cols: []map[int]bool{
				{22: true, 8: true, 21: true, 6: true, 1: true},
				{13: true, 2: true, 9: true, 10: true, 12: true},
				{17: true, 23: true, 14: true, 3: true, 20: true},
				{11: true, 4: true, 16: true, 18: true, 15: true},
				{0: true, 24: true, 7: true, 5: true, 19: true},
			},
		},
		{
			Numbers: [][]int{
				{3, 15, 0, 2, 22},
				{9, 18, 13, 17, 5},
				{19, 8, 7, 25, 23},
				{20, 11, 10, 24, 4},
				{14, 21, 16, 12, 6},
			},
			Rows: []map[int]bool{
				{3: true, 15: true, 0: true, 2: true, 22: true},
				{9: true, 18: true, 13: true, 17: true, 5: true},
				{19: true, 8: true, 7: true, 25: true, 23: true},
				{20: true, 11: true, 10: true, 24: true, 4: true},
				{14: true, 21: true, 16: true, 12: true, 6: true},
			},
			Cols: []map[int]bool{
				{9: true, 3: true, 19: true, 20: true, 14: true},
				{18: true, 15: true, 8: true, 11: true, 21: true},
				{13: true, 0: true, 7: true, 10: true, 16: true},
				{17: true, 2: true, 25: true, 24: true, 12: true},
				{5: true, 22: true, 23: true, 4: true, 6: true},
			},
		},
		{
			Numbers: [][]int{
				{14, 21, 17, 24, 4},
				{10, 16, 15, 9, 19},
				{18, 8, 23, 26, 20},
				{22, 11, 13, 6, 5},
				{2, 0, 12, 3, 7},
			},
			Rows: []map[int]bool{
				{14: true, 21: true, 17: true, 24: true, 4: true},
				{10: true, 16: true, 15: true, 9: true, 19: true},
				{18: true, 8: true, 23: true, 26: true, 20: true},
				{22: true, 11: true, 13: true, 6: true, 5: true},
				{2: true, 0: true, 12: true, 3: true, 7: true},
			},
			Cols: []map[int]bool{
				{14: true, 10: true, 18: true, 22: true, 2: true},
				{21: true, 16: true, 8: true, 11: true, 0: true},
				{17: true, 15: true, 23: true, 13: true, 12: true},
				{24: true, 9: true, 26: true, 6: true, 3: true},
				{4: true, 19: true, 20: true, 5: true, 7: true},
			},
		},
	},
}

func TestParts(t *testing.T) {
	util.RunTests(t, testData, []util.TestCase{
		{Desc: "Part 1", PartFunc: part1, Expected: 4512},
		{Desc: "Part 2", PartFunc: part2, Expected: 1924},
	})
}

func TestParseBoard(t *testing.T) {
	game := parseInput(testData)

	if !reflect.DeepEqual(game.DrawnNumbers, expected.DrawnNumbers) {
		t.Errorf("Expected drawn numbers: %v Got: %v", expected.DrawnNumbers, game.DrawnNumbers)
	}

	if len(game.Boards) != len(expected.Boards) {
		t.Errorf("Expected number of boards: %d Got: %d", len(expected.Boards), len(game.Boards))
	}

	for i := 0; i < len(game.Boards); i++ {
		if len(game.Boards[i].Numbers) != len(expected.Boards[i].Numbers) {
			t.Errorf("Expected number of rows for board %d: %d Got: %d",
				i, len(expected.Boards[i].Numbers), len(game.Boards[i].Numbers))
		}

		for r := 0; r < len(game.Boards[i].Numbers); r++ {
			if !reflect.DeepEqual(game.Boards[i].Numbers[r], expected.Boards[i].Numbers[r]) {
				t.Errorf("Expected board %d row %d: %v Got: %v", i, r,
					game.Boards[i].Numbers[r], expected.Boards[i].Numbers[r])
			}
		}

		for r := 0; r < len(game.Boards[i].Rows); r++ {
			if !reflect.DeepEqual(game.Boards[i].Rows[r], expected.Boards[i].Rows[r]) {
				t.Errorf("Expected board %d row map %d: %v Got: %v", i, r,
					game.Boards[i].Rows[r], expected.Boards[i].Rows[r])
			}
		}

		for c := 0; c < len(game.Boards[i].Cols); c++ {
			if !reflect.DeepEqual(game.Boards[i].Cols[c], expected.Boards[i].Cols[c]) {
				t.Errorf("Expected board %d column map %d: %v Got: %v", i, c,
					game.Boards[i].Cols[c], expected.Boards[i].Cols[c])
			}
		}
	}
}
