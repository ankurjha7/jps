package jps

import (
	"container/heap"
	"fmt"
	"math"
)

type aStarNode struct {
	node   Node
	gscore float64
	fscore float64
}

//Path struct to hold the shortest path result
type Path struct {
	Nodes  []Node
	Weight float64
}

// aStarQueue is an A* priority queue.
type aStarQueue struct {
	indexOf map[Node]int
	nodes   []aStarNode
}

func (q *aStarQueue) Less(i, j int) bool {
	return q.nodes[i].fscore < q.nodes[j].fscore
}

func (q *aStarQueue) Swap(i, j int) {
	q.indexOf[q.nodes[i].node] = j
	q.indexOf[q.nodes[j].node] = i
	q.nodes[i], q.nodes[j] = q.nodes[j], q.nodes[i]
}

func (q *aStarQueue) Len() int {
	return len(q.nodes)
}

func (q *aStarQueue) Push(x interface{}) {
	n := x.(aStarNode)
	q.indexOf[n.node] = len(q.nodes)
	q.nodes = append(q.nodes, n)
}

func (q *aStarQueue) Pop() interface{} {
	n := q.nodes[len(q.nodes)-1]
	q.nodes = q.nodes[:len(q.nodes)-1]
	delete(q.indexOf, n.node)
	return n
}

func (q *aStarQueue) update(node Node, g, f float64) {
	i, ok := q.indexOf[node]
	if !ok {
		return
	}
	q.nodes[i].gscore = g
	q.nodes[i].fscore = f
	heap.Fix(q, i)
}

func (q *aStarQueue) node(node *Node) (aStarNode, bool) {
	loc, ok := q.indexOf[*node]
	if ok {
		return q.nodes[loc], true
	}
	return aStarNode{}, false
}

//AStarWithJump astar implementation with neighbours filtered by jump point search .
func AStarWithJump(matrix [][]uint8, start Node, goal Node, hchoice int) (*Path, error) {

	cameFrom := make(map[Node]Node)
	visited := make(map[Node]bool)
	weightMap := make(map[Node]float64)

	open := &aStarQueue{indexOf: make(map[Node]int)}

	heap.Push(open, aStarNode{node: start, gscore: 0, fscore: heuristic(&start, &goal, 1)})
	for open.Len() != 0 {
		u := heap.Pop(open).(aStarNode)
		uNode := u.node

		if uNode.row == goal.row && uNode.col == goal.col {
			break
		}

		visited[uNode] = true
		to := identifySuccessors(uNode, cameFrom, matrix, goal)
		for _, v := range to {
			if _, exist := visited[v]; exist {
				continue
			}
			w := weight(&uNode, &v, 1)
			g := u.gscore + w
			if n, exist := open.node(&v); !exist {
				heap.Push(open, aStarNode{node: v, gscore: g, fscore: g + heuristic(&v, &goal, 1)})
			} else if g < n.gscore {
				open.update(v, g, g+heuristic(&v, &goal, 1))
			}
			cameFrom[v] = uNode
			weightMap[v] = w
		}
	}
	pathNodes := make([]Node, 0)

	if _, exist := cameFrom[goal]; !exist {
		return nil, fmt.Errorf("destination not reacheable from source")
	}
	cur := goal
	totalWeight := 0.0

	for {
		pathNodes = append(pathNodes, cur)
		weight, _ := weightMap[cur]
		totalWeight += weight
		next, exist := cameFrom[cur]
		if !exist {
			break
		}
		cur = next
	}

	//reversing to get path from start to end
	return &Path{getReverse(pathNodes), totalWeight}, nil

}

func getReverse(nodes []Node) []Node {
	length := len(nodes)
	reversedNodes := make([]Node, length)
	for i := range nodes {
		reversedNodes[i] = nodes[length-i-1]
	}
	return reversedNodes
}

func weight(a, b *Node, hchoice int) float64 {
	xDist := float64(a.row - b.row)
	yDist := float64(a.col - b.col)
	return math.Sqrt(math.Pow(xDist, 2) + math.Pow(yDist, 2))
}

func heuristic(a, b *Node, hchoice int) float64 {
	xDist := float64(a.row - b.row)
	yDist := float64(a.col - b.col)
	return math.Sqrt(math.Pow(xDist, 2) + math.Pow(yDist, 2))
}
