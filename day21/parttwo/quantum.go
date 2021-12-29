package parttwo

import (
	"fmt"
	"sort"
	"strings"
)

func PlayGame(p1Start, p2Start int, winScore int) [2]int {
	players := [2]*player{
		newPlayer(p1Start),
		newPlayer(p2Start),
	}
	return play(players, Player1, winScore)
}

type PlayerType int

const (
	Player1 PlayerType = iota
	Player2
)

var cachedPlays map[string][2]int = make(map[string][2]int)
var cachedDiceRolls map[int]int
var diceRolls []int

func init() {
	cachedDiceRolls = rollDice()
	for roll := range cachedDiceRolls {
		diceRolls = append(diceRolls, roll)
	}
	sort.Ints(diceRolls)
}

func rollDice() map[int]int {
	rolls := make(map[int]int)
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			for k := 1; k <= 3; k++ {
				rolls[i+j+k]++
			}
		}
	}
	return rolls
}

func play(ps [2]*player, turn PlayerType, winScore int) [2]int {
	cacheKey := createKey(ps, turn, winScore)
	if res, ok := cachedPlays[cacheKey]; ok {
		return res
	}

	results := [2]int{}
	for _, roll := range diceRolls {
		n := cachedDiceRolls[roll]
		players := [2]*player{
			ps[Player1].copy(),
			ps[Player2].copy(),
		}

		player := players[turn]
		player.move(roll)

		if player.won(winScore) {
			results[turn] += n
		} else {
			subresults := play(players, nextTurn(turn), winScore)
			results[Player1] += subresults[Player1] * n
			results[Player2] += subresults[Player2] * n
		}
	}

	cachedPlays[cacheKey] = results
	return results
}

func createKey(ps [2]*player, turn PlayerType, winScore int) string {
	var s strings.Builder
	s.WriteString("1:")
	s.WriteString(ps[Player1].cacheKey())
	s.WriteString(",2:")
	s.WriteString(ps[Player2].cacheKey())
	s.WriteString(",t")
	if turn == Player1 {
		s.WriteString("1")
	} else {
		s.WriteString("2")
	}
	s.WriteString(",w")
	s.WriteString(fmt.Sprintf("%d", winScore))
	return s.String()
}

func nextTurn(turn PlayerType) PlayerType {
	if turn == Player1 {
		return Player2
	}
	return Player1
}

type player struct {
	pos   int
	score int
}

func newPlayer(start int) *player {
	if start < 1 || start > 10 {
		panic("start must be between 1 and 10")
	}
	return &player{pos: start, score: 0}
}

func (p *player) copy() *player {
	return &player{p.pos, p.score}
}

func (p *player) won(winScore int) bool {
	return p.score >= winScore
}

func (p *player) move(n int) {
	p.pos += n
	if p.pos > 10 {
		p.pos = (p.pos-1)%10 + 1
	}
	p.score += p.pos
}

func (p *player) cacheKey() string {
	return fmt.Sprintf("p%ds%d", p.pos, p.score)
}
