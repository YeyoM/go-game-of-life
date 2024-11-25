package helpers

import (
	"github.com/gdamore/tcell/v2"
)

func RenderBoard(screen tcell.Screen, board Board, reservedSpaceX, reservedSpaceY int) {
	for i := 0; i < board.Height; i++ {
		for j := 0; j < board.Width; j++ {
			if board.Board[i][j] == true {
				screen.SetContent(j, i, '⬜', nil, tcell.StyleDefault.Background(tcell.ColorDefault))
			}
			if board.Board[i][j] == false {
				screen.SetContent(j, i, '⬛', nil, tcell.StyleDefault.Background(tcell.ColorDefault))
			}
		}
	}

}
