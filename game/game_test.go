package game

import (
	"fmt"
	"testing"
)

func TestGameMove1(t *testing.T) {
	b := Start()

	x, err := b.MakeMove(1, 0)
	x, err = x.MakeMove(2, 4)
	x, err = x.MakeMove(1, 8)
	//	x, err = x.MakeMove(2, 2)

	if err != nil {
		fmt.Println(err)
		return
	}

	//x, _ = x.MakeMove(1, 5)

	move, score := GetNextMove(*x)
	x.Print()

	fmt.Println("Move: ", move)
	fmt.Println("score: ", score)
	fmt.Println("player: ", x.PlayerToMove)
}
