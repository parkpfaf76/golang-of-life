package models

// grid origin from top-left (always select cell by row X col)
// cell grid
type CellGrid2D struct {
	numColumns int
	numRows int
	cells [][]*cell
}


func NewCellGrid2D(numRows int, numColumns int) CellGrid2D {
	cells := [numRows][numColumns]*cell{}

	for _, array := range cells {
        for j := range cellRow {
			cells = append(
				cells,
				newCell(newCellRules(new []int{2,3}, 3))) // hard code rules in each cell for now (2 or 3 neighbors to stay alive, 3 neighbors to spawn)
        }
    }

	g := cellGrid2D{numRows, numColumns, cells}
	return g
}

// retrive cell
func (c cellGrid2D) GetCellAtPos(rowIdx int, colIdx int) *cell {
	return g.cells[rowIdx][colIdx]
}
