// Package main is a console application for my first attempt at golang
package main

import (
	"fmt"
	"strconv"

	"github.com/philipf/tictactoe/game"
	"github.com/philipf/tictactoe/score"
)

func main() {
	clearScreen()
	board := game.Start()
	board.Print()

	for true {
		bestMove, currentScore, node := game.GetNextMove(*board)
		fmt.Printf("Best move: %d, score: %d\n", bestMove, currentScore)

		fmt.Printf("What is the next move for player %d:", board.PlayerToMove)
		var i string
		_, err := fmt.Scanf("%f\n", &i)

		var move int
		if i == "?" {
			node.Print(0)
			continue
		} else {
			move, err = strconv.Atoi(i)
		}

		if err != nil {
			fmt.Println(err)
			continue
		}

		board, err = board.MakeMove(board.PlayerToMove, move)
		clearScreen()
		board.Print()

		if err != nil {
			fmt.Println(err)
			continue
		}

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

func clearScreen() {
	// Needs work doesn't clear screens on all terminal windows
	print("\033[H\033[2J")
}
