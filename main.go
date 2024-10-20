package main

import (
	"collections/graph"
	"fmt"
)

func main() {
	g := graph.NewGraph[int]()
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 3)
	g.AddEdge(1, 4)
	g.AddEdge(2, 4)

	g.DFS(0, func(t int) {
		fmt.Printf("t: %v\n", t)
	})
}
