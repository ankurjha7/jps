package jps

type Node struct {
	row int
	col int
}

//GetCol returns column value for a node. in terms of x and y this is x
func (node *Node) GetCol() int {
	return node.col
}

//GetRow returns row value for a node . in terms of x and y this is y
func (node *Node) GetRow() int {
	return node.row
}

//GetNode returns node object for a row,columns value
func GetNode(row, col int) Node {
	return Node{row: row, col: col}
}

func (node *Node) equals(otherNode *Node) bool {
	return (node.row == otherNode.row) && (node.col == otherNode.col)
}

func blocked(row, col, dRow, dCol int, matrix [][]uint8) bool {
	if row+dRow < 0 || row+dRow >= len(matrix) {
		return true
	}
	if col+dCol < 0 || col+dCol >= len(matrix[0]) {
		return true
	}
	if matrix[row+dRow][col+dCol] == 1 {
		return true
	}
	return false
}

func dblock(row, col, dRow, dCol int, matrix [][]uint8) bool {
	if matrix[row-dRow][col] == 1 && matrix[row-dRow][col-dCol] == 1 {
		return true
	}
	return false
}

func direction(row, col, pRow, pCol int) (int, int) {
	dRow := getSign(row - pRow)
	dCol := getSign(col - pCol)
	if row-pRow == 0 {
		dRow = 0
	}
	if col-pCol == 0 {
		dCol = 0
	}
	return dRow, dCol
}

func getSign(num int) int {
	if num >= 0 {
		return 1
	}
	return -1
}

func getAllUnblockedNeighbours(row, col int, matrix [][]uint8) []*Node {
	neighbours := make([]*Node, 0)
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if !(i == 0 && j == 0) {
				if !blocked(row, col, i, j, matrix) {
					neighbours = append(neighbours, &Node{row: row + i, col: col + j})
				}
			}
		}
	}
	return neighbours
}

func nodeNeighbours(row, col int, parent *Node, matrix [][]uint8) []*Node {
	if parent == nil {
		return getAllUnblockedNeighbours(row, col, matrix) // returning all unblocked nodes for start node
	}
	neighbours := make([]*Node, 0)

	dRow, dCol := direction(row, col, parent.row, parent.col)

	if dRow != 0 && dCol != 0 {
		if !blocked(row, col, 0, dCol, matrix) {
			neighbours = append(neighbours, &Node{row, col + dCol})
		}
		if !blocked(row, col, dRow, 0, matrix) {
			neighbours = append(neighbours, &Node{row + dRow, col})
		}
		if (!blocked(row, col, 0, dCol, matrix) || !blocked(row, col, dRow, 0, matrix)) &&
			!blocked(row, col, dRow, dCol, matrix) {
			neighbours = append(neighbours, &Node{row + dRow, col + dCol})
		}
		if blocked(row, col, -dRow, 0, matrix) {
			neighbours = append(neighbours, &Node{row - dRow, col + dCol})
		}
		if blocked(row, col, 0, -dCol, matrix) {
			neighbours = append(neighbours, &Node{row + dRow, col - dCol})
		}
	} else {
		if dRow == 0 {
			if !blocked(row, col, dRow, 0, matrix) {
				if !blocked(row, col, 0, dCol, matrix) {
					neighbours = append(neighbours, &Node{row, col + dCol})
				}
				if blocked(row, col, 1, 0, matrix) {
					neighbours = append(neighbours, &Node{row + 1, col + dCol})
				}
				if blocked(row, col, -1, 0, matrix) {
					neighbours = append(neighbours, &Node{row - 1, col + dCol})
				}
			}
		} else {
			if !blocked(row, col, dRow, 0, matrix) {
				if !blocked(row, col, dRow, 0, matrix) {
					neighbours = append(neighbours, &Node{row + dRow, col})
				}
				if blocked(row, col, 0, 1, matrix) {
					neighbours = append(neighbours, &Node{row + dRow, col + 1})
				}
				if blocked(row, col, 0, -1, matrix) {
					neighbours = append(neighbours, &Node{row + dRow, col - 1})
				}
			}
		}
	}
	return neighbours
}

func jump(row, col, dRow, dCol int, matrix [][]uint8, goal *Node) *Node {

	nRow := row + dRow
	nCol := col + dCol
	if blocked(nRow, nCol, 0, 0, matrix) {
		return nil
	}

	if nRow == goal.row && nCol == goal.col {
		return &Node{row: nRow, col: nCol}
	}

	if dRow != 0 && dCol != 0 {
		// var1 := !blocked(nRow, nCol, -dRow, dCol, matrix)
		// var2 := blocked(nRow, nCol, -dRow, 0, matrix)
		// log.Printf("foudn it %s %s ", var1, var2)
		if (!blocked(nRow, nCol, -dRow, dCol, matrix) && blocked(nRow, nCol, -dRow, 0, matrix)) ||
			(!blocked(nRow, nCol, dRow, -dCol, matrix) && blocked(nRow, nCol, 0, -dCol, matrix)) {
			return &Node{row: nRow, col: nCol}
		}

		if jump(nRow, nCol, dRow, 0, matrix, goal) != nil || jump(nRow, nCol, 0, dCol, matrix, goal) != nil {
			return &Node{row: nRow, col: nCol}
		}

		if dblock(nRow, nCol, dRow, dCol, matrix) {
			return nil
		}
	} else {
		if dRow != 0 {
			if (!blocked(nRow, nCol, dRow, 1, matrix) && blocked(nRow, nCol, 0, 1, matrix)) ||
				(!blocked(nRow, nCol, dRow, -1, matrix) && blocked(nRow, nCol, 0, -1, matrix)) {
				return &Node{row: nRow, col: nCol}
			}

		} else {
			if (!blocked(nRow, nCol, 1, dCol, matrix) && blocked(nRow, nCol, 1, 0, matrix)) ||
				(!blocked(nRow, nCol, -1, dCol, matrix) && blocked(nRow, nCol, -1, 0, matrix)) {
				return &Node{row: nRow, col: nCol}
			}
		}
	}

	return jump(nRow, nCol, dRow, dCol, matrix, goal)
}

func identifySuccessors(current Node, parentNodeMap map[Node]Node, matrix [][]uint8, goal Node) []Node {
	successors := make([]Node, 0)
	var parent *Node
	if p, exist := parentNodeMap[current]; exist {
		parent = &p
	}
	neighbours := nodeNeighbours(current.row, current.col, parent, matrix)

	for _, cell := range neighbours {
		dRow := cell.row - current.row
		dCol := cell.col - current.col

		jumpPoint := jump(current.row, current.col, dRow, dCol, matrix, &goal)

		if jumpPoint != nil {
			successors = append(successors, *jumpPoint)
		}
	}
	return successors
}
