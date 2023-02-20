package models

// structure for single cell
type cell struct {
	numNeighbors int
	isAlive bool
	cellRules cellRules
}

func newCell(cellRules cellRules) *cell {
	c := cell{0, false, cellRules}
	return &c
}

// mutability methods
func (c *cell) SpawnCell() {
	c.isAlive = true
}

func (c *cell) KillCell() {
	c.isAlive = false
}

func (c *cell) DecrementNumNeighbors() {
	c.numNeighbors--
}

func (c *cell) IncrementNumNeighbors() {
	c.numNeighbors++
}

func (c *cell) UpdateNumNeighbors(numNeighbors int) {
	c.numNeighbors = numNeighbors
}

// return true if cell state change
func (c *cell) UpdateToNextCellState() bool {

	if c.isAlive {
		for _, currVal := range c.cellRules.numNeighborsToStayAlive {
			if currVal == c.numNeighbors {
				return false
			}
		}
		c.KillCell()
		return true
	} else {
		if c.cellRules.numNeighborsToSpawn == c.numNeighbors {
			c.SpawnCell()
			return true
		}
	}

	return false
}