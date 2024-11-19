package helpers

func GetTerminalMiddlePoint(size Size) (c Coordinate) {
	// calculate the middle point
	c.x = size.Width / 2
	c.y = size.Height / 2

	return c
}
