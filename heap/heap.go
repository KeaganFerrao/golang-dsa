package heap

import (
	"fmt"
)

// A Heap is implemented as an array because its always a complete binary tree
// without any holes. All the levels are full except the lowest which fills always left to right
// For any index i,
// left child index => 2*i + 1
// right child index => 2*i + 2
// parent index => (i - 1) / 2
type Heap struct {
	arr []int
}

func (h *Heap) Insert(key int) {
	h.arr = append(h.arr, key)
	h.heapifyUp(len(h.arr) - 1)
}

// ExtractMin removes and returns the smallest element (root)
func (h *Heap) ExtractMin() (int, error) {
	if len(h.arr) == 0 {
		return 0, fmt.Errorf("heap is empty")
	}

	min := h.arr[0]
	h.arr[0] = h.arr[len(h.arr)-1]
	h.arr = h.arr[:len(h.arr)-1]
	h.heapifyDown(0)

	return min, nil
}

// Peek returns the smallest element without removing it
func (h *Heap) Peek() (int, error) {
	if len(h.arr) == 0 {
		return 0, fmt.Errorf("heap is empty")
	}
	return h.arr[0], nil
}

func (h *Heap) heapifyUp(index int) {
	for index > 0 {
		parent := (index - 1) / 2
		if h.arr[index] < h.arr[parent] {
			// Swap
			h.arr[index], h.arr[parent] = h.arr[parent], h.arr[index]

			// Continue moving up the heap
			index = parent
		} else {
			// Everything is in proper order so break
			break
		}
	}
}

func (h *Heap) heapifyDown(index int) {
	lastIndex := len(h.arr) - 1
	for {
		leftChild := 2*index + 1
		rightChild := 2*index + 2
		smallest := index

		if leftChild <= lastIndex && h.arr[leftChild] < h.arr[smallest] {
			smallest = leftChild
		}

		if rightChild <= lastIndex && h.arr[rightChild] < h.arr[smallest] {
			smallest = rightChild
		}

		if smallest != index {
			h.arr[index], h.arr[smallest] = h.arr[smallest], h.arr[index]
			index = smallest
		} else {
			break
		}
	}
}
