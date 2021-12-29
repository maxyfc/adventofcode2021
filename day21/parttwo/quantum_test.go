package parttwo

import "testing"

func TestPlayerMove(t *testing.T) {
	p := newPlayer(1)

	tests := []struct {
		move, score, pos int
	}{
		{10, 1, 1},
		{9, 11, 10},
		{1, 12, 1},
	}

	for i, test := range tests {
		p.move(test.move)
		if p.score != test.score {
			t.Fatalf("Move: %d Expected Score: %d Got: %d", i, p.score, test.score)
		}
		if p.pos != test.pos {
			t.Fatalf("Move: %d Expected Pos: %d Got: %d", i, p.pos, test.pos)
		}
	}
}
