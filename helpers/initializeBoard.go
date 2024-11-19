package helpers

import "math/rand/v2"

func InitializeBoard(size Size) Board {
	var board Board
	board.Width = size.Width
	board.Height = size.Height
	board.Board = make([][]bool, board.Height)

	for i := range board.Board {
		board.Board[i] = make([]bool, board.Width)
	}

	for i := 0; i < board.Height; i++ {
		for j := 0; j < board.Width; j++ {
			if j%2 == 0 {
				if rand.IntN(100) < 70 {
					board.Board[i][j] = false
				} else {
					board.Board[i][j] = true
				}
			} else {
				board.Board[i][j] = false
			}
		}
	}

	return board
}
