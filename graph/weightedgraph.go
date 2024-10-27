package graph

import (
	"collections/queue"
	"log"
)

// Edge represents an edge in the WeightedGraph
type Edge[T comparable] struct {
	to     T
	weight int
}

// WeightedGraph represents a weighted, undirected WeightedGraph using an adjacency list
type WeightedGraph[T comparable] struct {
	vertices map[T][]Edge[T]
}

// NewGraph initializes a WeightedGraph with the given number of vertices
func NewWeightedGraph[T comparable]() *WeightedGraph[T] {
	return &WeightedGraph[T]{
		vertices: make(map[T][]Edge[T]),
	}
}

// AddEdge adds a weighted undirected edge to the WeightedGraph
func (g *WeightedGraph[T]) AddEdge(from T, to T, weight int) {
	g.vertices[from] = append(g.vertices[from], Edge[T]{to, weight})
	g.vertices[to] = append(g.vertices[to], Edge[T]{from, weight})
}

// PrimMST calculates the Minimum Spanning Tree (MST) using Prim's algorithm
func (g *WeightedGraph[T]) PrimMST() int {
	// Minimum total weight of edges
	totalWeight := 0
	// Map to keep a track of all visited nodes
	visited := make(map[T]bool)

	pq := queue.NewPriorityQueue[T]()

	// Start from any random vertex
	// Since a map is unordered, we are looping over it once to get a random element
	// Enqueue this vertex
	for key := range g.vertices {
		pq.Enqueue(queue.Item[T]{Item: key, Priority: 0})
		break
	}

	// We loop untill all the nodes are not visited
	for pq.Length() > 0 {
		// Dequeue the vertex with the least weight, since its a priority queue(Heap)
		// the dequeue will give the node with the next least weight
		// Time: O(logN), since we use a heap internally
		item, error := pq.Dequeue()
		if error != nil {
			log.Fatal("Something went wrong", error)
		}
		vertex := item.Item
		weight := item.Priority

		// If that is already visited, just skip
		if visited[vertex] {
			continue
		}

		// Mark the vertex as visited and add its weight to the total
		visited[vertex] = true
		totalWeight += weight

		// Add all adjacent edges to the priority queue if not already visited
		for _, edge := range g.vertices[vertex] {
			if !visited[edge.to] {
				// This would automatically rearrange nodes in decreasing order of weights, since its a priority queue
				// So the next time we can just do a normal dequeue, and it will give the element with
				// the least weight which is needed in prims algorithm
				pq.Enqueue(queue.Item[T]{Item: edge.to, Priority: edge.weight})
			}
		}
	}

	return totalWeight
}
