package models

type gameState struct {
	cellGrid2D cellGrid2D
	totalNumberAlive := 0
	totalNumberDeadCells := width * height
}

// Entry point for controller to act on cells
func NewGameState(gameBoardWidth int, gameBoardHeight int) gameState {
	grid := cellGrid2D[gameBoardWidth][gameBoardHeight]
	totalNumberAlive := 0
	totalNumberDeadCells := width * height
	newGame := game{grid}
	return newGame
}