package helpers

import (
	"github.com/gdamore/tcell/v2"
)

func GetTerminalSize(screen tcell.Screen) (size Size) {
	w, h := screen.Size()
	size.Width = w
	size.Height = h

	return size
}
