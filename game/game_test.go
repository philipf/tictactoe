package game

import (
	"fmt"
	"testing"
)

func TestGameMove1(t *testing.T) {
	b := Start()

	//x, err := b.MakeMove(1, 4)
	// x, err = x.MakeMove(2, 1)
	// x, err = x.MakeMove(1, 0)
	// x, err = x.MakeMove(2, 8)
	// x, err = x.MakeMove(1, 3)

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	//x, _ = x.MakeMove(1, 5)

	move, score := GetNextMove(b)
	b.Print()

	fmt.Println("Move: ", move)
	fmt.Println("score: ", score)
	fmt.Println("player: ", b.PlayerToMove)
}
