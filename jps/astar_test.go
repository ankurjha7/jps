package jps

import "testing"

/* Here inclusion of forced neighbour is essential to find shortest path
and the forced neghbour is against the direction of travel */
func TestForcedNeighbour1(t *testing.T) {
	matrix := [][]uint8{
		[]uint8{0, 1, 0, 0, 0}, // here node {0,2} is a forced neighbour for  {1,1}
		[]uint8{0, 0, 1, 0, 0},
		[]uint8{1, 1, 1, 1, 0},
		[]uint8{0, 0, 0, 0, 0},
	}
	path, err := AStarWithJump(matrix, GetNode(3, 2), GetNode(0, 0), 1)
	if err != nil {
		t.Errorf("error while creating path")
	}
	if len(path.Nodes) == 0 {
		t.Errorf("invalid path")
	}

}

/* Here inclusion of forced neighbour is essential to find shortest path
and the forced neghbour is against the direction of travel */
func TestForcedNeighbour2(t *testing.T) {
	matrix := [][]uint8{
		[]uint8{0, 1, 1, 1, 1}, // here node {0,2} is a forced neighbour for  {1,1}
		[]uint8{0, 0, 1, 0, 0},
		[]uint8{0, 1, 1, 1, 0},
		[]uint8{0, 0, 0, 0, 0},
	}
	path, err := AStarWithJump(matrix, GetNode(0, 0), GetNode(3, 2), 1)
	if err != nil {
		t.Errorf("error while creating path")
	}
	if len(path.Nodes) == 0 {
		t.Errorf("invalid path")
	}

}

/* Here inclusion of forced neighbour is essential to find shortest path
and the forced neghbour is against the direction of travel */
func TestForcedNeighbour3(t *testing.T) {
	matrix := [][]uint8{
		[]uint8{0, 1, 1, 1, 1}, // here node {0,2} is a forced neighbour for  {1,1}
		[]uint8{1, 0, 1, 0, 0},
		[]uint8{0, 1, 1, 1, 0},
		[]uint8{0, 0, 0, 0, 0},
	}
	path, err := AStarWithJump(matrix, GetNode(0, 0), GetNode(3, 2), 1)
	if err != nil {
		t.Errorf("error while creating path")
	}
	if len(path.Nodes) == 0 {
		t.Errorf("invalid path")
	}

}

/* Shortest path without blocker */
func TestForcedNeighbour4(t *testing.T) {
	matrix := [][]uint8{
		[]uint8{0, 0, 0, 0, 0}, // here node {0,2} is a forced neighbour for  {1,1}
		[]uint8{0, 0, 0, 0, 0},
		[]uint8{0, 0, 0, 0, 0},
		[]uint8{0, 0, 0, 0, 0},
	}
	path, err := AStarWithJump(matrix, GetNode(0, 0), GetNode(3, 2), 1)
	if err != nil {
		t.Errorf("error while creating path")
	}
	if len(path.Nodes) == 0 {
		t.Errorf("invalid path")
	}

}
