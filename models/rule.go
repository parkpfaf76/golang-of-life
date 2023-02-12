package models

// each cell will have rules associated
type rules struct {
	numNeighborsToStayAlive []int
	numNeighborsToSpawn int
}

func newRules(numNeighborsToStayAlive []int, numNeighborsToSpawn int) rules {
	r := rules{numNeighborsToStayAlive, numNeighborsToSpawn}
	return r
}