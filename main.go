package main

import (
	"github.com/YeyoM/game-of-life/helpers"
	"github.com/gdamore/tcell/v2"
	"github.com/micmonay/keybd_event"

	"os"
)

func main() {

	arguments := os.Args

	//  Init screen
	screen := helpers.InitScreen()
	defer screen.Fini()

	// Get terminal size
	var terminalSize helpers.Size
	terminalSize = helpers.GetTerminalSize(screen)

	//var middlePoint helpers.Coordinate
	//middlePoint = helpers.GetTerminalMiddlePoint(terminalSize)

	var currentSize helpers.Size
	currentSize = terminalSize

	// Init view board
	var currentBoard helpers.Board
	var nextBoard helpers.Board

	// Random Initialization
	// currentBoard = helpers.InitializeBoard(currentSize)

	// Sample Pattern Initialization
	// currentBoard = helpers.InitializeBoardPattern1(currentSize)

	// glider
	// currentBoard = helpers.InitializeBoardGlider(currentSize)

	// custom pattern (from file)
	file := arguments[1]
	if file == "" {
		panic("Please provide a file")
	}
	currentBoard = helpers.InitializeBoardCustom(currentSize, file)

	// custom pattern (from file)

	// Main game loop
	for {
		screen.Show() // show screen

		// Display information
		helpers.RenderBoard(screen, currentBoard, 0, 0)

		// Press enter to update board
		kb, err := keybd_event.NewKeyBonding()
		if err != nil {
			panic(err)
		}

		kb.SetKeys(keybd_event.VK_ENTER)
		kb.Launching()

		event := screen.PollEvent() // get event from screen

		switch event := event.(type) {
		case *tcell.EventResize:
			screen.Sync()
			currentSize = helpers.GetTerminalSize(screen)

		case *tcell.EventKey:
			if event.Key() == tcell.KeyEscape || event.Key() == tcell.KeyCtrlC {
				screen.Clear()
				screen.Fini()
				return
			}

			if event.Key() == tcell.KeyEnter {
				nextBoard = helpers.UpdateBoard(currentBoard)
				currentBoard = nextBoard
				nextBoard.Board = currentBoard.Board
			}

		default:
			// Do nothing
		}

		helpers.Sleep(50)
	}
}
