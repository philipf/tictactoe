package board

import "fmt"

const boardSize = 9

// Board represents a tic-tac-to board
type Board struct {
	State        [boardSize]int
	PlayerToMove int
	LastMove     int
}

// New tic-tac-toe board
func New() *Board {
	board := Board{PlayerToMove: 1, LastMove: -1}
	return &board
}

// MakeMove to set the move
func (board *Board) MakeMove(player int, move int) (*Board, error) {
	if player != board.PlayerToMove {
		return board, fmt.Errorf("Invalid player %d, expected %d", player, board.PlayerToMove)
	}

	if move < 0 || move >= boardSize {
		return board, fmt.Errorf("Invalid move %d, should be between 0 and %d", move, boardSize)
	}

	if board.State[move] != 0 {
		return board, fmt.Errorf("Invalid move, block %d already contains a move %d", move, board.State[move])
	}

	changedBoard := board.clone()

	if player == 1 {
		changedBoard.PlayerToMove = 2
	} else {
		changedBoard.PlayerToMove = 1
	}

	changedBoard.LastMove = move
	changedBoard.State[move] = player

	return &changedBoard, nil
}

func (board *Board) reset() {
	for index := 0; index < boardSize; index++ {
		board.State[index] = 0
	}
}

// Print the board layout
func (board Board) Print() {
	s := board.State

	text := fmt.Sprintf("+---+---+---+\n"+
		"| %s | %s | %s |\n"+
		"+---+---+---+\n"+
		"| %s | %s | %s |\n"+
		"+---+---+---+\n"+
		"| %s | %s | %s |\n"+
		"+---+---+---+",
		translate(s[0]), translate(s[1]), translate(s[2]),
		translate(s[3]), translate(s[4]), translate(s[5]),
		translate(s[6]), translate(s[7]), translate(s[8]))

	fmt.Println(text)
}

func translate(val int) (result string) {
	result = "-"

	if val == 1 {
		result = "O"
	} else if val == 2 {
		result = "X"
	}

	return
}

func (board Board) clone() Board {
	clone := Board{
		State:        board.State,
		PlayerToMove: board.PlayerToMove,
		LastMove:     board.LastMove}
	return clone
}
