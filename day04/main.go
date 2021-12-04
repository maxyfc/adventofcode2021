package main

import (
	"adventofcode2021/pkg/strutil"
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var inputData string

func main() {
	fmt.Printf("Part 1: %d\n", part1(inputData))
	fmt.Printf("Part 2: %d\n", part2(inputData))
}

func part1(input string) int {
	game := parseInput(input)
	for _, n := range game.DrawnNumbers {
		for _, b := range game.Boards {
			b.Remove(n)

			if b.Won {
				return b.RemainingSum() * n
			}
		}
	}

	return 0
}

func part2(input string) int {
	game := parseInput(input)

	boards := make(map[*Board]bool)
	for _, b := range game.Boards {
		boards[b] = true
	}

	for _, n := range game.DrawnNumbers {
		var won []*Board
		for b := range boards {
			b.Remove(n)

			if b.Won {
				if len(boards) == 1 { // last remaining board
					return b.RemainingSum() * n
				}

				won = append(won, b)
			}
		}

		for _, b := range won {
			delete(boards, b)
		}
	}

	return 0
}

type BingoGame struct {
	DrawnNumbers []int
	Boards       []*Board
}

type Board struct {
	Numbers [][]int
	Rows    []map[int]bool
	Cols    []map[int]bool
	Won     bool
}

func (b *Board) Init() {
	b.Cols = make([]map[int]bool, 0, len(b.Numbers))
	for range b.Numbers[0] {
		b.Cols = append(b.Cols, make(map[int]bool))
	}
	for _, r := range b.Numbers {
		row := make(map[int]bool)
		for ci, num := range r {
			row[num] = true
			b.Cols[ci][num] = true
		}
		b.Rows = append(b.Rows, row)
	}
}

func (b *Board) Remove(n int) {
	if b.Won {
		return
	}

	for _, r := range b.Rows {
		_, ok := r[n]
		if ok {
			delete(r, n)
		}
		if len(r) == 0 {
			b.Won = true
		}
	}

	for _, c := range b.Cols {
		_, ok := c[n]
		if ok {
			delete(c, n)
		}
		if len(c) == 0 {
			b.Won = true
		}
	}
}

func (b *Board) RemainingSum() int {
	sum := 0
	for _, row := range b.Rows {
		for n := range row {
			sum += n
		}
	}
	return sum
}

func parseInput(input string) BingoGame {
	s := bufio.NewScanner(strings.NewReader(input))
	s.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		// Copied from bufio.ScanLines with some modifications

		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}

		// We need find 2 new lines next to each other with the possibility of
		// carriage return in the middle
		offset := 0
		for offset < len(data) {
			i := bytes.IndexByte(data[offset:], '\n')
			j := bytes.IndexByte(data[offset+i+1:], '\n')
			if i >= 0 && j >= 0 && j <= 1 {
				// We have a full newline-terminated line.
				return offset + i + j + 2, dropCR(data[0 : offset+i]), nil
			} else if i < 0 || j < 0 {
				break
			}
			offset += i + j + 2
		}

		// If we're at EOF, we have a final, non-terminated line. Return it.
		if atEOF {
			return len(data), dropCR(data), nil
		}
		// Request more data.
		return 0, nil, nil
	})

	if !s.Scan() {
		panic("Missing drawn numbers line. Should be the first line in the input.")
	}

	game := BingoGame{}

	// parse the drawn numbers
	nums := strings.Split(s.Text(), ",")
	game.DrawnNumbers = strutil.MustAtoiSlice(nums)

	// parse the boards
	for s.Scan() {
		boardInput := s.Text()
		rows := strutil.SplitLines(boardInput)
		board := Board{}
		for _, row := range rows {
			r := strutil.MustAtoiSlice(strings.Fields(row))
			board.Numbers = append(board.Numbers, r)
		}

		if len(board.Numbers) != 5 {
			panic(fmt.Sprintf("Error parsing board data. Expected 5 rows but got %d.",
				len(board.Numbers)))
		}

		board.Init()

		game.Boards = append(game.Boards, &board)
	}

	if err := s.Err(); err != nil {
		panic(fmt.Sprintf("Error parsing board, %v", err))
	}

	return game
}

// Remove the CR. Coped from go stdlib source code.
func dropCR(data []byte) []byte {
	if len(data) > 0 && data[len(data)-1] == '\r' {
		return data[0 : len(data)-1]
	}
	return data
}
