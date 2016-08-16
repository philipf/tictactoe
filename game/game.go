package game

import (
	"fmt"
	"ttt/board"
	"ttt/minimax"
	"ttt/score"
)

// Start new game
func Start() board.Board {
	b := board.New()
	return *b
}

// GetNextMove return the next move and the score
func GetNextMove(board board.Board) (int, int) {
	rootNode := minimax.New()
	suggestBestMove(board, &rootNode, board.PlayerToMove)
	//rootNode.Print(0)
	rootNode.Evaluate()

	move := rootNode.GetBestChildNode().Data.(int)
	score := *rootNode.Score
	return move, score
}

func suggestBestMove(board board.Board, node *minimax.Node, playerToMove int) {
	for m := 0; m < len(board.State); m++ {
		if board.State[m] != 0 {
			continue
		}

		nextBoard, err := board.MakeMove(board.PlayerToMove, m)

		if err != nil {
			fmt.Println(err)
			return
		}

		nextScore := score.Score(*nextBoard, playerToMove)

		if nextScore == -2 {
			nextNode := node.Add(m)
			suggestBestMove(*nextBoard, nextNode, playerToMove)
		} else {
			node.AddTerminal(nextScore, m)
		}
	}
}
