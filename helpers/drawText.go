package helpers

import (
	"github.com/gdamore/tcell/v2"
)

func DrawText(screen tcell.Screen, x int, y int, text string) {
	for i, c := range text {
		screen.SetContent(x+i, y, c, nil, tcell.StyleDefault)
	}
}
