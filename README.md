# jps
Golang implementation of Jump point search . Jump point search minimizes A* execution time by jumping nodes which won't be contributing in optimizing path. Detailed explaination available at - https://users.cecs.anu.edu.au/~dharabor/data/papers/harabor-grastien-icaps14.pdf

## Usage 
```go
package main

import (
	"fmt"

	"github.com/ankurjha7/jps"
)

func main() {

	grid := [][]uint8{
		{0, 1, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{1, 1, 1, 1, 0},
		{0, 0, 0, 0, 0},
	}
	start := jps.GetNode(0, 0)
	end := jps.GetNode(3, 2)
	path, err := jps.AStarWithJump(grid, start, end, 1)
	if err == nil {
		fmt.Printf("Path is : ")
		for _, node := range path.Nodes {
			fmt.Printf("%d %d -> ", node.GetRow(), node.GetCol())
		}
		fmt.Printf("\nTotal distance %f", path.Weight)
	}

}


