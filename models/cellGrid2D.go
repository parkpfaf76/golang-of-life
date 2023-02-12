package models

// cell grid
type cellGrid2D struct {
	width int
	height int
	cells [][]cell
}

func newCellGrid2D(width int, height int) cellGrid2D {
	cells := [width][height]cell
	g := cellGrid2D{width, height, cells, totalNumberAlive, totalNumberDeadCells}
	return g
}

// retrive cell
func (c cellGrid2D) getCell(posX int, posY int) cell {
	return g.cells[posX][posY]
}
