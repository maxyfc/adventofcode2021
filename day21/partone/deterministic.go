// This solution onlys for part 1. Need a new solution for part 2.
package partone

func PlayGame(p1Start, p2Start int, d Dice, winningScore int) (bool, int) {
	p1 := NewPlayer(p1Start, d)
	p2 := NewPlayer(p2Start, d)
	for {
		p1.RollAndMove()
		if p1.Score() >= winningScore {
			return true, p2.Score() * d.NumRoll()
		}

		p2.RollAndMove()
		if p2.Score() >= winningScore {
			return false, p1.Score() * d.NumRoll()
		}
	}
}

type Dice interface {
	Roll() int
	NumRoll() int
}

type DeterministicDice struct {
	curr    int
	numRoll int
}

func NewDeterministicDice() *DeterministicDice {
	return &DeterministicDice{
		curr: 0,
	}
}

func (d *DeterministicDice) Roll() int {
	d.curr++
	d.numRoll++
	if d.curr > 100 {
		d.curr = 1
	}
	return d.curr
}

func (d *DeterministicDice) NumRoll() int {
	return d.numRoll
}

type Player struct {
	pos   int
	dice  Dice
	score int
}

func NewPlayer(startPos int, d Dice) *Player {
	if startPos < 1 || startPos > 10 {
		panic("startPos must be between 1 and 10")
	}
	return &Player{
		pos:   startPos,
		dice:  d,
		score: 0,
	}
}

func (p *Player) RollAndMove() {
	for i := 0; i < 3; i++ {
		r := p.dice.Roll()
		p.move(r)
	}
	p.score += p.pos
}

func (p *Player) move(n int) {
	p.pos += n
	if p.pos > 10 {
		p.pos = (p.pos-1)%10 + 1
	}
}

func (p *Player) Score() int {
	return p.score
}

func (p *Player) Pos() int {
	return p.pos
}
