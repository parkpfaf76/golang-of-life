package models

// grid origin from top-left (always select cell by row X col)
// cell grid
type CellGrid2D struct {
	numColumns int
	numRows int
	cells [][]*cell
}


func NewCellGrid2D(numRows int, numColumns int) CellGrid2D {
	cells := make([][]*cell,numRows, numColumns)

	for _, cellRow := range cells {
        for j := range cellRow {
			cellRow[j] = newCell(newCellRules([]int{2,3}, 3)) // hard code rules in each cell for now (2 or 3 neighbors to stay alive, 3 neighbors to spawn)
        }
    }

	g := CellGrid2D{numRows, numColumns, cells}
	return g
}

// retrive cell
func (cg CellGrid2D) GetCellAtPos(rowIdx int, colIdx int) *cell {
	return cg.cells[rowIdx][colIdx]
}
