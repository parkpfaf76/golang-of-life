package models

// structure for single cell
type cell struct {
	numNeighbors int
	isAlive bool
	cellRules cellRules
}

func newCell(cellRules cellRules) cell {
	c := cell{0, false, cellRules}
	return c
}

// mutability methods
func (c *cell) spawnCell() {
	c.isAlive = true
}

func (c *cell) killCell() {
	c.isAlive = false
}

func (c *cell) decrementNumNeighbors() {
	c.numNeighbors--
}

func (c *cell) incrementNumNeighbors() {
	c.numNeighbors++
}

func (c *cell) updateNumNeighbors(numNeighbors int) {
	c.numNeighbors = numNeighbors
}

// return true if cell state change
func (c *cell) shouldCellUpdate() bool {

	if c.isAlive {
		for _, currVal := range c.cellRules.numNeighborsToStayAlive {
			if currVal == c.numNeighbors {
				return false
			}
		}
		return true
	} else {
		if c.cellRules.numNeighborsToSpawn == c.numNeighbors {
			return true
		}
	}

	return false
}