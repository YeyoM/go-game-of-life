package helpers

func UpdateBoard(prevBoard Board) (newBoard Board) {
	// Write header for this update

	newBoard.Width = prevBoard.Width
	newBoard.Height = prevBoard.Height
	newBoard.Board = make([][]bool, newBoard.Height)
	for i := range newBoard.Board {
		newBoard.Board[i] = make([]bool, newBoard.Width)
	}
	for i := 0; i < prevBoard.Height; i++ {
		for j := 0; j < prevBoard.Width; j++ {
			neighbours := checkNeighbours(prevBoard, j, i)

			if prevBoard.Board[i][j] {
				if neighbours < 2 {
					newBoard.Board[i][j] = false
				} else if neighbours == 2 || neighbours == 3 {
					newBoard.Board[i][j] = true
				} else if neighbours > 3 {
					newBoard.Board[i][j] = false
				}
			} else {
				if neighbours == 3 {
					newBoard.Board[i][j] = true
				}
			}
		}
	}
	return newBoard
}

func checkNeighbours(board Board, x int, y int) (neighbours int) {
	neighbours = 0
	// There are 8 possible neighbours
	neighboursPositions := []Coordinate{
		{x - 2, y - 1},
		{x - 2, y},
		{x - 2, y + 1},
		{x, y - 1},
		{x, y + 1},
		{x + 2, y - 1},
		{x + 2, y},
		{x + 2, y + 1},
	}
	for _, neighbour := range neighboursPositions {
		if neighbour.x >= 0 && neighbour.x < board.Width && neighbour.y >= 0 && neighbour.y < board.Height {
			if board.Board[neighbour.y][neighbour.x] {
				neighbours++
			}
		} else {
			continue
		}
	}
	return neighbours
}
