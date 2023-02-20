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
	c.numNeightbors--
}

func (c *cell) IncrementNumNeighbors() {
	c.numNeightbors++
}

func (c *cell) UpdateNumNeighbors(numNeightbors int) {
	c.numNeightbors = numNeightbors
}

// return true if cell state change
func (c *cell) UpdateToNextCellState() bool {

	if c.isAlive {
		for _, currVal := range c.rules.numNeighborsToStayAlive {
			if currVal == c.numNeighbors {
				break
			}
		}
		c.killCell()
	} else {
		if c.cellRules.numNeighborsToSpawn == c.numNeighbors {
			c.spawnCell()
		}
	}
}