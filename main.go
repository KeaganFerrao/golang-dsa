package main

import (
	"github.com/KeaganFerrao/golang-dsa/graph"
	"fmt"
)

func main() {
	g := graph.NewGraph[int]()
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(3, 4)
	g.AddEdge(2, 2)

	hasCycle := g.ContinsCycle(1)
	fmt.Printf("hasCycle: %v\n", hasCycle)
}
