package models

// structure for single cell
type cell struct {
	numNeighbors int
	isAlive bool
	cellRules cellRules
}

func newCell(cellRules cellRules) *cell {
	c := cell{0, isAlive, cellRules}
	return &c
}

// mutability methods
func (c *cell) spawnCell() {
	c.isAlive = true
}

func (c *cell) killCell() {
	c.isAlive = false
}

func (c *cell) decrementNumNeighbors() {
	c.numNeightbors--
}

func (c *cell) incrementNumNeighbors() {
	c.numNeightbors++
}

func (c *cell) updateNumNeighbors(numNeightbors int) {
	c.numNeightbors = numNeightbors
}

// return true if cell state change
func (c *cell) updateToNextCellState() bool {

	if c.IsAlive() {
		for _, currVal := range c.rules.numNeighborsToStayAlive {
			if currVal == c.numNeighbors {
				break
			}
		}
		c.killCell()
	} 
	else {
		if c.cellRules.numNeighborsToSpawn == c.numNeighbors {
			c.spawnCell()
		}
	}
}