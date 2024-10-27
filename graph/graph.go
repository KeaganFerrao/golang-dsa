package graph

import (
	"collections/queue"
	"collections/stack"
	"log"
)

// Adjacency list representation of an undirected graph
type Graph[T comparable] struct {
	vertices map[T][]T
}

func NewGraph[T comparable]() *Graph[T] {
	return &Graph[T]{vertices: make(map[T][]T)}
}

// An undirected graph
func (g *Graph[T]) AddEdge(v T, w T) {
	g.vertices[v] = append(g.vertices[v], w)
	g.vertices[w] = append(g.vertices[w], v)
}

func (g *Graph[T]) BFS(start T, cb func(T)) {
	visited := make(map[T]bool)
	queue := queue.NewQueue[T]()

	visited[start] = true
	queue.Enqueue(start)

	for queue.Length() > 0 {
		element, _ := queue.Dequeue()

		cb(*element)

		for _, neighbor := range g.vertices[*element] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue.Enqueue(neighbor)
			}
		}
	}
}

func (g *Graph[T]) DFS(start T, cb func(T)) {
	visited := make(map[T]bool)
	stack := stack.NewStack[T]()

	stack.Push(start)

	for stack.Length() > 0 {
		elementRef, _ := stack.Pop()
		element := *elementRef

		if !visited[element] {
			visited[element] = true
			cb(element)

			// Push all unvisited neighbors
			for _, neighbor := range g.vertices[element] {
				if !visited[neighbor] {
					stack.Push(neighbor)
				}
			}
		}
	}
}

// We do a DFS to check for cycles in a undirected graph
func (g *Graph[T]) ContinsCycle(start T) bool {
	visited := make(map[T]bool)
	parent := make(map[T]*T)
	stack := stack.NewStack[T]()

	stack.Push(start)
	parent[start] = nil

	for stack.Length() > 0 {
		elementRef, error := stack.Pop()
		if error != nil {
			log.Fatal(error)
		}
		element := *elementRef

		if !visited[element] {
			visited[element] = true
		}

		// Push all unvisited neighbors
		for _, neighbor := range g.vertices[element] {
			if !visited[neighbor] {
				parent[neighbor] = &element
				stack.Push(neighbor)
			} else if parent[element] == nil || neighbor != *parent[element] {
				// If the neighbor is visited and is not the parent of the current element,
				// then we found a cycle
				return true
			}
		}
	}

	return false
}
