package models

// each cell will have rules associated
type cellRules struct {
	numNeighborsToStayAlive []int
	numNeighborsToSpawn int
}

func newCellRules(numNeighborsToStayAlive []int, numNeighborsToSpawn int) cellRules {
	r := cellRules{numNeighborsToStayAlive, numNeighborsToSpawn}
	return r
}