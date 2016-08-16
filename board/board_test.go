package board

import "testing"

func TestNew(t *testing.T) {
	b := New()

	b.Print()

	if b == nil {
		t.Error("board is nil")
	}

	if len(b.State) != 9 {
		t.Error("board state is not 9: ", len(b.State))
	}

	for i := 0; i < 9; i++ {
		if b.State[i] != 0 {
			t.Errorf("Invalid state %d for %d", b.State[i], i)
		}
	}
}

func TestClone(t *testing.T) {
	b1 := New()
	b1.State[0] = 1
	b1.State[8] = 2

	b2 := b1.clone()
	b1.reset()

	b2.Print()

	if b2.State[0] != 1 {
		t.Error("clone failed for [0]", b2.State[0])
	}

	if b2.State[8] != 2 {
		t.Error("clone failed for [8]", b2.State[8])
	}
}

func TestMakeMoveInvalidPlayer(t *testing.T) {
	b1 := New()

	_, err := b1.MakeMove(2, 0)

	if err == nil {
		t.Error("Expected invalid player")
	}
}

func TestMakeMoveInvalidMove(t *testing.T) {
	b1 := New()

	_, err := b1.MakeMove(1, 9)

	if err == nil {
		t.Error("Expected invalid move")
	}
}

func TestMakeMoveInvalidMoveAlreadyDone(t *testing.T) {
	b1 := New()

	b2, _ := b1.MakeMove(1, 0)
	_, err := b2.MakeMove(2, 0)

	if err == nil {
		t.Error("Expected invalid move: ", err)
	}
}
