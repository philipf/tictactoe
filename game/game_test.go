package game

import (
	"fmt"
	"testing"

	"github.com/philipf/tictactoe/board"
)

func TestGameInital(t *testing.T) {
	var x *board.Board
	var err interface{}

	x = Start()

	if err != nil {
		fmt.Println(err)
		return
	}

	move, score, _ := GetNextMove(*x)
	x.Print()

	fmt.Println("Move: ", move)
	fmt.Println("score: ", score)
	fmt.Println("player: ", x.PlayerToMove)

	if score != 0 {
		t.Error("Should reach stalemate 0, actual: ", score)
	}

	if x.PlayerToMove != 1 {
		t.Error("Expected player 1, ", x.PlayerToMove)
	}
}

func TestGameMove1(t *testing.T) {
	var x *board.Board
	var err interface{}

	x = Start()

	x, err = x.MakeMove(1, 4)

	if err != nil {
		fmt.Println(err)
		return
	}

	move, score, _ := GetNextMove(*x)
	x.Print()

	fmt.Println("Move: ", move)
	fmt.Println("score: ", score)
	fmt.Println("player: ", x.PlayerToMove)

	if score != 0 {
		t.Error("Score should be equal 0, actual: ", score)
	}

	if x.PlayerToMove != 2 {
		t.Error("Expected player 2, ", x.PlayerToMove)
	}
}

func TestGameMovePlayer2Mistake(t *testing.T) {
	var x *board.Board
	var err interface{}

	x = Start()

	x, err = x.MakeMove(1, 4)
	x, err = x.MakeMove(2, 1)

	if err != nil {
		fmt.Println(err)
		return
	}

	move, score, _ := GetNextMove(*x)
	x.Print()

	fmt.Println("Move: ", move)
	fmt.Println("score: ", score)
	fmt.Println("player: ", x.PlayerToMove)

	if score != 1 {
		t.Error("Score should favour player 1 and be 2, actual: ", score)
	}

	if x.PlayerToMove != 1 {
		t.Error("Expected player 1, ", x.PlayerToMove)
	}
}
