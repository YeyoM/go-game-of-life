package helpers

import (
	"bufio"
	"os"
)

func InitializeBoardCustom(size Size, fileName string) Board {
	// Initialize the board with given size
	var board Board
	board.Width = size.Width
	board.Height = size.Height
	board.Board = make([][]bool, board.Height)
	for i := range board.Board {
		board.Board[i] = make([]bool, board.Width)
	}

	// initialize the board with all false
	for i := 0; i < board.Height; i++ {
		for j := 0; j < board.Width; j++ {
			board.Board[i][j] = false
		}
	}

	// Open and read the file
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	lineNumber := 0
	columnNumber := 0

	// Starting position at (6,6)
	startX := 20
	startY := 80

	for scanner.Scan() {
		line := scanner.Text()
		// Process each character in the line
		for _, char := range line {
			if char == 'O' {
				board.Board[startX+lineNumber][startY+columnNumber] = true
			}
			columnNumber += 2
		}
		columnNumber = 0
		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return board
}

