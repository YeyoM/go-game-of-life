package helpers

func InitializeBoardGlider(size Size) Board {
	var board Board
	board.Width = size.Width
	board.Height = size.Height
	board.Board = make([][]bool, board.Height)

	for i := range board.Board {
		board.Board[i] = make([]bool, board.Width)
	}

	for i := 0; i < board.Height; i++ {
		for j := 0; j < board.Width; j++ {
			board.Board[i][j] = false
		}
	}

	board.Board[6][6] = true
	board.Board[7][6] = true
	board.Board[8][6] = true
	board.Board[8][4] = true
	board.Board[7][2] = true

	return board
}
