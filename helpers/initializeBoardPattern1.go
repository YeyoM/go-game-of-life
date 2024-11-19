package helpers

func InitializeBoardPattern1(size Size) Board {
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

	board.Board[2][2] = true
	board.Board[2][4] = true
	board.Board[2][6] = true

	return board
}
