package models

type GameState struct {
	cellGrid2D cellGrid2D
	totalNumberAlive int
	totalNumberDeadCells int
}

func NewGameState(numRows int, numColumns int) *GameState {
	grid := newCellGrid2D(numRows, numColumns)
	totalNumberAlive := 0
	totalNumberDeadCells := numRows * numColumns
	newGameState := GameState{grid, totalNumberAlive, totalNumberDeadCells}
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

// point to brand new game state
func (gs *GameState) ClearCurrentGameState() {
	gs = NewGameState(gs.cellGrid2D.numRows, gs.cellGrid2D.numColumns)
}

// UpdateToNextFrame
func (gs *GameState) UpdateToNextFrame()  {
	var didCellUpdate bool
	cells := gs.cellGrid2D.getCells()

	for rowIdx := 0; rowIdx < gs.cellGrid2D.numRows; rowIdx++ {
        for colIdx := 0; colIdx < gs.cellGrid2D.numColumns; colIdx++ {
			didCellUpdate = cells[rowIdx][colIdx].updateToNextCellState();

			if(didCellUpdate) {
				gs.handleCellUpdate(rowIdx, colIdx)
			}
        }
    }
}

func (gs *GameState) SpawnCell(rowIdx int, colIdx int)  {
	cells := gs.cellGrid2D.getCells()
	cells[rowIdx][colIdx].spawnCell()
	gs.handleCellSpawnsEvent(rowIdx, colIdx)
}

func (gs *GameState) KillCell(rowIdx int, colIdx int)  {
	cells := gs.cellGrid2D.getCells()
	cells[rowIdx][colIdx].killCell()
	gs.handleCellDiesEvent(rowIdx, colIdx)
}

func (gs *GameState) handleCellUpdate(rowIdx int, colIdx int)  {
	cells := gs.cellGrid2D.getCells()

	if cells[rowIdx][colIdx].isAlive {
		gs.handleCellSpawnsEvent(rowIdx, colIdx)
	} else {
		gs.handleCellDiesEvent(rowIdx, colIdx)
	}
}

// helpers
func (gs *GameState) incrementSurroundingNeighbors(rowIdx int, colIdx int)  {
	var startRowIdx, endRowIdx, startColIdx, endColIdx = gs.getColRowSurroundingIndicies(rowIdx, colIdx)

	for i := startRowIdx; i <= endRowIdx; i++ {
        for j := startColIdx; j <= endColIdx; j++ {
			if(i == rowIdx && j == colIdx) {
				continue;
			}

			cells := gs.cellGrid2D.getCells()
			cells[i][j].incrementNumNeighbors()
		}
    }
}

func (gs *GameState) decrementSurroundingNeighbors(rowIdx int, colIdx int)  {
	var startRowIdx, endRowIdx, startColIdx, endColIdx = gs.getColRowSurroundingIndicies(rowIdx, colIdx)

	for i := startRowIdx; i <= endRowIdx; i++ {
        for j := startColIdx; j <= endColIdx; j++ {
			if(i == rowIdx && j == colIdx) {
				continue;
			}

			cells := gs.cellGrid2D.getCells()
			cells[i][j].decrementNumNeighbors()
		}
    }
}

func (gs *GameState) handleCellDiesEvent(rowIdx int, colIdx int)  {
	gs.totalNumberAlive--;
	gs.totalNumberDeadCells++;
	gs.decrementSurroundingNeighbors(rowIdx, colIdx)
}

func (gs *GameState) handleCellSpawnsEvent(rowIdx int, colIdx int)  {
	gs.totalNumberAlive++;
	gs.totalNumberDeadCells--;
	gs.incrementSurroundingNeighbors(rowIdx, colIdx)
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