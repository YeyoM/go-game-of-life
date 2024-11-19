package helpers

type Board struct {
	Width  int
	Height int

	Board [][]bool
}

type Size struct {
	Width  int
	Height int
}

type Coordinate struct {
	x int
	y int
}
