package score

import "ttt/board"

var winningLines = [8][3]int{
	{0, 1, 2}, // 0
	{3, 4, 5}, // 1
	{6, 7, 8}, // 2
	{0, 3, 6}, // 3
	{1, 4, 7}, // 4
	{2, 5, 8}, // 5
	{0, 4, 8}, // 6
	{2, 4, 6}} // 7

// Score a board
func Score(board board.Board, playerToMove int) int {
	bs := board.State

	for line := 0; line < len(winningLines); line++ {
		score := scoreLine(bs, winningLines[line], playerToMove)

		if score != 0 {
			return score
		}
	}

	for index := 0; index < len(bs); index++ {
		if bs[index] == 0 {
			return -2 // not completed
		}
	}

	return 0 // stalemate
}

func scoreLine(bs [9]int, line [3]int, playerToMove int) int {
	var cell1 = bs[line[0]]
	var cell2 = bs[line[1]]
	var cell3 = bs[line[2]]

	if cell1 == 0 || cell2 == 0 || cell3 == 0 {
		return 0
	}

	// a win: X X X
	if cell1 == cell2 && cell2 == cell3 {
		if cell1 == playerToMove {
			return 1
		} else {
			return -1
		}
	}

	return 0
}
