package controller

import "github.com/parkpfaf76/golang-of-life/view"

// each cell will have rules associated
type controller struct {
	gameState gameState
	view view
}

func NewController(gameState gameState, view view) controller {
	
}