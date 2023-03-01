package models

type GameState struct {
	cellGrid2D cellGrid2D
	totalNumberAlive int
	totalNumberDeadCells int
}

func NewGameState(numRows int, numColumns int) *GameState {
	currentGrid := newCellGrid2D(numRows, numColumns)
	totalNumberAlive := 0
	totalNumberDeadCells := numRows * numColumns
	newGameState := GameState{currentGrid, totalNumberAlive, totalNumberDeadCells}
	return &newGameState
}

// GetCurrentGameStateGrid - no need to expose cell - will return bool[][] true=alive false=dead
func (gs *GameState) GetCurrentGameStateGrid() [][]bool {
	stateArray := make([][]bool, gs.cellGrid2D.numRows)

	for rowIdx := range stateArray {
		stateArray[rowIdx] = make([]bool, gs.cellGrid2D.numColumns)
        for colIdx := range stateArray[rowIdx] {
			currcells := gs.cellGrid2D.getCells()
			stateArray[rowIdx][colIdx] = currcells[rowIdx][colIdx].isAlive
        }
    }

	return stateArray
}

func (gs *GameState) ClearCurrentGameState() {
	gs.cellGrid2D = newCellGrid2D(gs.cellGrid2D.numRows, gs.cellGrid2D.numColumns)
}

// UpdateToNextFrame
func (gs *GameState) UpdateToNextFrame()  {
	var shouldCellUpdate bool
	cells := gs.cellGrid2D.getCells()

	cellsCopy := make([][]cell, len(cells))
	for i := range cells {
		cellsCopy[i] = make([]cell, len(cells[i]))
		copy(cellsCopy[i], cells[i])
	}

	for rowIdx := 0; rowIdx < gs.cellGrid2D.numRows; rowIdx++ {
        for colIdx := 0; colIdx < gs.cellGrid2D.numColumns; colIdx++ {
			shouldCellUpdate = cells[rowIdx][colIdx].shouldCellUpdate();

			if(shouldCellUpdate) {
				cellsCopy = gs.handleCellUpdate(rowIdx, colIdx, cellsCopy)
			}
        }
    }

	gs.cellGrid2D.cells = cellsCopy
}

func (gs *GameState) SpawnCell(rowIdx int, colIdx int)  {
	cells := gs.cellGrid2D.getCells()
	cells[rowIdx][colIdx].spawnCell()
	gs.cellGrid2D.cells = gs.handleCellSpawnsEvent(rowIdx, colIdx, cells)
}

func (gs *GameState) KillCell(rowIdx int, colIdx int)  {
	cells := gs.cellGrid2D.getCells()
	cells[rowIdx][colIdx].killCell()
	gs.cellGrid2D.cells = gs.handleCellDiesEvent(rowIdx, colIdx, cells)
}

func (gs *GameState) handleCellUpdate(rowIdx int, colIdx int, cells [][]cell) [][]cell  {
	if cells[rowIdx][colIdx].isAlive {
		cells[rowIdx][colIdx].killCell()
		return gs.handleCellDiesEvent(rowIdx, colIdx, cells)	
	} else {
		cells[rowIdx][colIdx].spawnCell()
		return gs.handleCellSpawnsEvent(rowIdx, colIdx, cells)
	}
}

// helpers
func (gs *GameState) incrementSurroundingNeighbors(rowIdx int, colIdx int, cells [][]cell) [][]cell {
	var startRowIdx, endRowIdx, startColIdx, endColIdx = gs.getColRowSurroundingIndicies(rowIdx, colIdx)

	for i := startRowIdx; i <= endRowIdx; i++ {
        for j := startColIdx; j <= endColIdx; j++ {
			if(i == rowIdx && j == colIdx) {
				continue;
			}

			cells[i][j].incrementNumNeighbors()
		}
    }

	return cells
}

func (gs *GameState) decrementSurroundingNeighbors(rowIdx int, colIdx int, cells [][]cell) [][]cell {
	var startRowIdx, endRowIdx, startColIdx, endColIdx = gs.getColRowSurroundingIndicies(rowIdx, colIdx)

	for i := startRowIdx; i <= endRowIdx; i++ {
        for j := startColIdx; j <= endColIdx; j++ {
			if(i == rowIdx && j == colIdx) {
				continue;
			}

			cells[i][j].decrementNumNeighbors()
		}
    }

	return cells
}

func (gs *GameState) handleCellDiesEvent(rowIdx int, colIdx int, cells [][]cell) [][]cell {
	gs.totalNumberAlive--;
	gs.totalNumberDeadCells++;
	return gs.decrementSurroundingNeighbors(rowIdx, colIdx, cells)
}

func (gs *GameState) handleCellSpawnsEvent(rowIdx int, colIdx int, cells [][]cell) [][]cell {
	gs.totalNumberAlive++;
	gs.totalNumberDeadCells--;
	return gs.incrementSurroundingNeighbors(rowIdx, colIdx, cells)
}

// helper to get neighbor row/col indicies given a position => returns rowStartIdx, rowEndIdx, colStartIdx, colEndIdx
func (gs *GameState) getColRowSurroundingIndicies(rowIdx int, colIdx int) (int, int, int, int) {
	var startRowIdx, endRowIdx, startColIdx, endColIdx int

	if rowIdx - 1 < 0 {
		startRowIdx = 0
	} else {
		startRowIdx = rowIdx - 1
	}

	if rowIdx + 1 >= gs.cellGrid2D.numRows {
		endRowIdx = gs.cellGrid2D.numRows - 1
	} else {
		endRowIdx = rowIdx + 1
	}

	if colIdx - 1 < 0 {
		startColIdx = 0
	} else {
		startColIdx = colIdx - 1
	}

	if colIdx + 1 >= gs.cellGrid2D.numColumns {
		endColIdx = gs.cellGrid2D.numColumns - 1
	} else {
		endColIdx = colIdx + 1
	}

    return startRowIdx, endRowIdx, startColIdx, endColIdx
}