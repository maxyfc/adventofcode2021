package partone

import (
	"testing"
)

func TestDeterministicDiceRoll(t *testing.T) {
	d := NewDeterministicDice()
	for i := 1; i <= 100; i++ {
		val := d.Roll()
		if i != val {
			t.Fatalf("Expected roll: %d Got %d", i, val)
		}
	}
	if d.NumRoll() != 100 {
		t.Fatalf("Expected number of 100 rolls. Got: %d", d.NumRoll())
	}
	val := d.Roll()
	if val != 1 {
		t.Fatalf("Expected to wrap back to 1. Got: %d", val)
	}
	if d.NumRoll() != 101 {
		t.Fatalf("Expected number of 101 rolls. Got: %d", d.NumRoll())
	}
}

func TestPlayerRollAndMove(t *testing.T) {
	d := NewDeterministicDice()
	p1 := NewPlayer(4, d)
	p2 := NewPlayer(8, d)

	steps := []struct {
		p1Pos   int
		p1Score int
		p2Pos   int
		p2Score int
	}{
		{10, 10, 3, 3},
		{4, 14, 6, 9},
		{6, 20, 7, 16},
		{6, 26, 6, 22},
	}

	for _, step := range steps {
		p1.RollAndMove()
		if p1.Pos() != step.p1Pos {
			t.Errorf("Expected p1 pos: %d Got: %d", step.p1Pos, p1.Pos())
		}
		if p1.Score() != step.p1Score {
			t.Errorf("Expected p1 score: %d Got: %d", step.p1Score, p1.Score())
		}

		p2.RollAndMove()
		if p2.Pos() != step.p2Pos {
			t.Errorf("Expected p2 pos: %d Got: %d", step.p2Pos, p2.Pos())
		}
		if p2.Score() != step.p2Score {
			t.Errorf("Expected p2 score: %d Got: %d", step.p2Score, p2.Score())
		}
	}
}
