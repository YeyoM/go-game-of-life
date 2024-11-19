package helpers

import (
	"github.com/gdamore/tcell/v2"
	"strconv"
)

// Display the current status of the game
// current terminal size
// middle point of the terminal

func DisplayInformation(screen tcell.Screen, terminalSize Size, middlePoint Coordinate, boardCoordinate bool) {
	DrawText(screen, 0, 0, "Terminal size: ")
	DrawText(screen, 0, 1, "Width: ")
	DrawText(screen, 0, 2, "Height: ")
	DrawText(screen, 0, 3, "Middle point: ")
	DrawText(screen, 0, 4, "X: ")
	DrawText(screen, 0, 5, "Y: ")
	DrawText(screen, 0, 6, "Value: ")

	DrawText(screen, 18, 1, strconv.Itoa(terminalSize.Width))
	DrawText(screen, 18, 2, strconv.Itoa(terminalSize.Height))
	DrawText(screen, 18, 4, strconv.Itoa(middlePoint.x))
	DrawText(screen, 18, 5, strconv.Itoa(middlePoint.y))
	DrawText(screen, 18, 6, strconv.FormatBool(boardCoordinate))
}
