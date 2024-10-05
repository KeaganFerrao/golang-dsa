package main

import (
	"fmt"
	"practice/heap"
)

func main() {
	heap := heap.Heap{}

	heap.Insert(11)
	heap.Insert(-99)
	heap.Insert(0)
	heap.Insert(1)
	heap.Insert(15)
	heap.Insert(9)
	heap.Insert(56)
	heap.Insert(6)

	fmt.Printf("heap: %v\n", heap)
}
