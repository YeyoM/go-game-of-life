package helpers

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"os"
)

func RenderBoard(screen tcell.Screen, board Board, reservedSpaceX, reservedSpaceY int) {
	// Create or open debug file
	debugFile, err := os.OpenFile("debug_board.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("Error opening debug file: %v\n", err)
		return
	}
	defer debugFile.Close()

	// Write header for this update
	fmt.Fprintf(debugFile, "\n--- New Board Render ---\n")

	for i := 0; i < board.Height; i++ {
		for j := 0; j < board.Width; j++ {
			if board.Board[i][j] == true {
				screen.SetContent(j, i, '⬜', nil, tcell.StyleDefault.Background(tcell.ColorDefault))
			}
			if board.Board[i][j] == false {
				screen.SetContent(j, i, '⬛', nil, tcell.StyleDefault.Background(tcell.ColorDefault))
			}
			if i >= 1 && j >= 1 && i <= 3 && j <= 3 {
				fmt.Fprintf(debugFile, "Coordinate (%d, %d) - State: %t\n", j, i, board.Board[i][j])
			}
		}
	}

}
