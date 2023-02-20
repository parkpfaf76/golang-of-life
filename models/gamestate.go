package models

import (
    "math"
)

type gameState struct {
	cellGrid2D cellGrid2D
	totalNumberAlive int
	totalNumberDeadCells int
}

// Entry point for controller to play game
func NewGameState(numRows int, numColumns int) *gameState {
	grid := newCellGrid2D(numRows, numColumns)
	totalNumberAlive := 0
	totalNumberDeadCells := numRows * numColumns
	newGameState := gameState{grid, totalNumberAlive, totalNumberDeadCells}
	return &newGameState
}

// Controller exposed methods
// GetCurrentGameStateGrid - no need to expose cell to controller - will return bool[][] true=alive false=dead
func (gs *gameState) GetCurrentGameStateGrid() [][]bool {
	stateArray := [gs.grid.numRows][gs.grid.numColumns]bool{}

	for rowIdx, _ := range gs.cellGrid2D {
        for colIdx, currCell := range cellRow {
			stateArray[rowIdx][colIdx] = currCell.isAlive;
        }
    }

	return stateArray
}

// UpdateToNextFrame
func (gs *gameState) UpdateToNextFrame()  {
	var didCellUpdate bool

	for rowIdx, cellRow := range gs.cellGrid2D {
        for colIdx, currCell := range cellRow {
			didCellUpdate = currCell.UpdateToNextCellState();

			if(didCellUpdate) {
				gs.handleCellUpdate(rowIdx, colIdx)
			}
        }
    }
}

func (gs *gameState) SpawnCell(rowIdx int, colIdx int)  {
	gs.cellGrid2D.getCellAtPos(rowIdx, colIdx).SpawnCell()
	gs.handleCellSpawnsEvent()
}

func (gs *gameState) KillCell(rowIdx int, colIdx int)  {
	
	gs.cellGrid2D.getCellAtPos(rowIdx, colIdx).KillCell()
	gs.handleCellDiesEvent()
}

func (gs *gameState) handleCellUpdate(rowIdx int, colIdx int)  {
	if c.isAlive() {
		gs.handleCellSpawnsEvent(rowIdx, colIdx)
	} else {
		gs.handleCellDiesEvent(rowIdx, colIdx)
	}
}

// helpers
func (gs *gameState) incrementSurroundingNeighbors(rowIdx int, colIdx int)  {
	var startRowIdx, endRowIdx, startColIdx, endColIdx = getColRowSurroundingIndicies(rowIdx, colIdx)

	for i := startRowIdx; i < endRowIdx; i++ {
        for j := startColIdx; j < endColIdx; i++ {
			if(i == rowIdx && j == colIdx) {
				continue;
			}

			gs.cellGrid2D.getCellAtPos(i, j).IncrementNumNeighbors()
		}
    }
}

func (gs *gameState) decrementSurroundingNeighbors(rowIdx int, colIdx int)  {
	var startRowIdx, endRowIdx, startColIdx, endColIdx = getColRowSurroundingIndicies(rowIdx, colIdx)

	for i := startRowIdx; i < endRowIdx; i++ {
        for j := startColIdx; j < endColIdx; i++ {
			if(i == rowIdx && j == colIdx) {
				continue;
			}

			gs.cellGrid2D.getCellAtPos(i, j).DecrementNumNeighbors()
		}
    }
}

func (gs *gameState) handleCellDiesEvent()  {
	gs.totalNumberAlive--;
	gs.totalNumberDeadCells++;
	gs.DecrementSurroundingNeighbors(rowIdx, colIdx)
}

func (gs *gameState) handleCellSpawnsEvent()  {
	gs.IncrementSurroundingNeighbors(rowIdx, colIdx)
	gs.totalNumberAlive++;
	gs.totalNumberDeadCells--;
}

// helper to get neighbor row/col indicies given a position => returns rowStartIdx, rowEndIdx, colStartIdx, colEndIdx
func (gs *gameState) getColRowSurroundingIndicies(rowIdx int, colIdx int) (int, int, int, int) {
	var startRowIdx, endRowIdx, startColIdx, endColIdx int

	startRowIdx = Min(0, rowIdx - 1)
	endRowIdx = Max(rowIdx + 1, gs.cellGrid2D.height)
	startColIdx = Min(colIdx - 1)
	endColIdx = Max(colIdx + 1)

    return startRowIdx, endRowIdx, startColIdx, endColIdx
}