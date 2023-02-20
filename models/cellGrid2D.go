package models

// grid origin from top-left (always select cell by row X col)
// cell grid
type cellGrid2D struct {
	numColumns, numRows int
	cells [][]*cell
}


func newCellGrid2D(numRows int, numColumns int) cellGrid2D {
	cells := make([][]*cell,numRows, numColumns)

	for _, cellRow := range cells {
        for j := range cellRow {
			cellRow[j] = newCell(newCellRules([]int{2,3}, 3)) // hard code rules in each cell for now (2 or 3 neighbors to stay alive, 3 neighbors to spawn)
        }
    }

	g := cellGrid2D{numRows, numColumns, cells}
	return g
}

// retrive cell
func (cg cellGrid2D) getCellAtPos(rowIdx int, colIdx int) *cell {
	return cg.cells[rowIdx][colIdx]
}
