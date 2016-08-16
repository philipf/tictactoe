package score

import (
	"testing"
	"ttt/board"
)

func TestIncomplete(t *testing.T) {
	b := board.New()

	score := Score(*b, 1)

	if score != -2 {
		t.Error("Expected incomplete -2, but score was: ", score)
	}
}

func TestPlayer1Win(t *testing.T) {
	b := board.New()

	b, _ = b.MakeMove(1, 0)
	b, _ = b.MakeMove(2, 3)
	b, _ = b.MakeMove(1, 1)
	b, _ = b.MakeMove(2, 4)
	b, _ = b.MakeMove(1, 2)

	score := Score(*b, 1)

	if score != 2 {
		t.Error("Expected player 1 win, but score was: ", score)
	}
}

func TestPlayer2Win(t *testing.T) {
	b := board.New()

	b, _ = b.MakeMove(1, 0)
	b, _ = b.MakeMove(2, 3)
	b, _ = b.MakeMove(1, 1)
	b, _ = b.MakeMove(2, 4)
	b, _ = b.MakeMove(1, 7)
	b, _ = b.MakeMove(2, 5)

	score := Score(*b, 2)

	if score != 2 {
		t.Error("Expected player 2 win, but score was: ", score)
	}
}

func TestStalemate(t *testing.T) {
	b := board.New()

	b.State[0] = 1
	b.State[1] = 2
	b.State[2] = 1
	b.State[3] = 2
	b.State[4] = 1
	b.State[5] = 2
	b.State[6] = 2
	b.State[7] = 1
	b.State[8] = 2

	b.Print()

	score := Score(*b, 1)

	if score != 0 {
		t.Error("Expected stalemate: ", score)
	}
}
