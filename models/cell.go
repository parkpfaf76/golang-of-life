package models

// structure for single cell
type cell struct {
	posX int
	poxY int
	numNeighbors int
	isAlive bool
	rules rules
}

func newCell(posX int, poxY int, rules rules) *cell {
	c := cell{posX, poxY, 0, isAlive, rules}
	return &c
}

// mutability methods
func (c *cell) spawnCell() {
	c.isAlive = true
}

func (c *cell) killCell() {
	c.isAlive = false
}

func (c *cell) updateNumNeighbors(numNeightbors int) {
	c.numNeightbors = numNeightbors
}

// get cell status
func (c cell) isAlive() bool {
	return c.isAlive
}

func (c cell) getRules() rules {
	return c.rules
}
