package models

type gameState struct {
	cellGrid2D cellGrid2D
	totalNumberAlive int
	totalNumberDeadCells int
}

func NewGameState(numRows int, numColumns int) *gameState {
	grid := newCellGrid2D(numRows, numColumns)
	totalNumberAlive := 0
	totalNumberDeadCells := numRows * numColumns
	newGameState := gameState{grid, totalNumberAlive, totalNumberDeadCells}
	return &newGameState
}

// GetCurrentGameStateGrid - no need to expose cell - will return bool[][] true=alive false=dead
func (gs *gameState) GetCurrentGameStateGrid() [][]bool {
	stateArray := make([][]bool, gs.cellGrid2D.numRows)

	for rowIdx, cellRow := range stateArray {
		stateArray[rowIdx] = make([]bool, gs.cellGrid2D.numColumns)
        for colIdx := range cellRow {
			stateArray[colIdx] = append(stateArray[colIdx], gs.cellGrid2D.getCellAtPos(rowIdx, colIdx).isAlive)
        }
    }

	return stateArray
}

// UpdateToNextFrame
func (gs *gameState) UpdateToNextFrame()  {
	var didCellUpdate bool

	for rowIdx := 0; rowIdx < gs.cellGrid2D.numRows; rowIdx++ {
        for colIdx := 0; colIdx < gs.cellGrid2D.numColumns; colIdx++ {
			didCellUpdate = gs.cellGrid2D.getCellAtPos(rowIdx, colIdx).updateToNextCellState();

			if(didCellUpdate) {
				gs.handleCellUpdate(rowIdx, colIdx)
			}
        }
    }
}

func (gs *gameState) SpawnCell(rowIdx int, colIdx int)  {
	gs.cellGrid2D.getCellAtPos(rowIdx, colIdx).spawnCell()
	gs.handleCellSpawnsEvent(rowIdx, colIdx)
}

func (gs *gameState) KillCell(rowIdx int, colIdx int)  {
	
	gs.cellGrid2D.getCellAtPos(rowIdx, colIdx).killCell()
	gs.handleCellDiesEvent(rowIdx, colIdx)
}

func (gs *gameState) handleCellUpdate(rowIdx int, colIdx int)  {
	cell := gs.cellGrid2D.getCellAtPos(rowIdx, colIdx)

	if cell.isAlive {
		gs.handleCellSpawnsEvent(rowIdx, colIdx)
	} else {
		gs.handleCellDiesEvent(rowIdx, colIdx)
	}
}

// helpers
func (gs *gameState) incrementSurroundingNeighbors(rowIdx int, colIdx int)  {
	var startRowIdx, endRowIdx, startColIdx, endColIdx = gs.getColRowSurroundingIndicies(rowIdx, colIdx)

	for i := startRowIdx; i < endRowIdx; i++ {
        for j := startColIdx; j < endColIdx; i++ {
			if(i == rowIdx && j == colIdx) {
				continue;
			}

			gs.cellGrid2D.getCellAtPos(i, j).incrementNumNeighbors()
		}
    }
}

func (gs *gameState) decrementSurroundingNeighbors(rowIdx int, colIdx int)  {
	var startRowIdx, endRowIdx, startColIdx, endColIdx = gs.getColRowSurroundingIndicies(rowIdx, colIdx)

	for i := startRowIdx; i < endRowIdx; i++ {
        for j := startColIdx; j < endColIdx; i++ {
			if(i == rowIdx && j == colIdx) {
				continue;
			}

			gs.cellGrid2D.getCellAtPos(i, j).decrementNumNeighbors()
		}
    }
}

func (gs *gameState) handleCellDiesEvent(rowIdx int, colIdx int)  {
	gs.totalNumberAlive--;
	gs.totalNumberDeadCells++;
	gs.decrementSurroundingNeighbors(rowIdx, colIdx)
}

func (gs *gameState) handleCellSpawnsEvent(rowIdx int, colIdx int)  {
	gs.incrementSurroundingNeighbors(rowIdx, colIdx)
	gs.totalNumberAlive++;
	gs.totalNumberDeadCells--;
}

// helper to get neighbor row/col indicies given a position => returns rowStartIdx, rowEndIdx, colStartIdx, colEndIdx
func (gs *gameState) getColRowSurroundingIndicies(rowIdx int, colIdx int) (int, int, int, int) {
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