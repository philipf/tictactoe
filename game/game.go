package game

import (
	"ttt/board"
	"ttt/minimax"
	"ttt/score"
)

// Start new game
func Start() *board.Board {
	b := board.New()
	return b
}

// GetNextMove return the next move and the score
func GetNextMove(b board.Board) (int, int, minimax.Node) {
	rootNode := minimax.New()
	calculatePossibleMoves(b, &rootNode, b.PlayerToMove)
	//	rootNode.Print(0)
	rootNode.Evaluate()

	node := *rootNode.GetBestChildNode()
	bestBoard := node.Data.(*board.Board)

	move := bestBoard.LastMove
	score := *rootNode.Score
	return move, score, rootNode
}

func calculatePossibleMoves(board board.Board, node *minimax.Node, playerToMove int) {
	for m := 0; m < len(board.State); m++ {
		if board.State[m] != 0 {
			continue
		}

		nextBoard, err := board.MakeMove(board.PlayerToMove, m)

		if err != nil {
			panic(err)
		}

		nextScore := score.Score(*nextBoard, playerToMove)

		if nextScore == -2 {
			nextNode := node.Add(nextBoard)
			calculatePossibleMoves(*nextBoard, nextNode, playerToMove)
		} else {
			node.AddTerminal(nextScore, nextBoard)
		}
	}
}
