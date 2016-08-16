package main

import (
	"fmt"
	"ttt/game"
	"ttt/score"
)

func main() {
	print("\033[H\033[2J")
	board := game.Start()

	board.Print()

	for true {
		move, currentScore := game.GetNextMove(*board)
		fmt.Printf("Best move: %d, score: %d\n", move, currentScore)

		var i int
		fmt.Printf("What is the next move for player %d:", board.PlayerToMove)
		_, err := fmt.Scan(&i)

		if err != nil {
			fmt.Println(err)
			continue
		}

		board, err = board.MakeMove(board.PlayerToMove, i)

		if err != nil {
			fmt.Println(err)
			continue
		}

		print("\033[H\033[2J")
		board.Print()

		gameScore := score.Score(*board, board.PlayerToMove)

		if gameScore == -2 {
			// game in progress
		} else if gameScore == 0 {
			fmt.Println("Game over: Draw")
			break
		} else {
			fmt.Printf("Game over: Player %d wins\n", board.PlayerToMove)
			break
		}
	}
}
