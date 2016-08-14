package board

// NewBoard creates a new board
func NewBoard() Board {
	b := Board{123}
	return b
}

// Board represents a tic-tac-to board
type Board struct {
	Score int
}
