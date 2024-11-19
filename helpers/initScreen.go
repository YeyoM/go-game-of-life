package helpers

import (
	"github.com/gdamore/tcell/v2"
	"os"
)

func InitScreen() tcell.Screen {
	screen, err := tcell.NewScreen()
	if err != nil {
		os.Exit(1)
	}
	if err := screen.Init(); err != nil {
		os.Exit(1)
	}
	screen.Clear()
	return screen
}
