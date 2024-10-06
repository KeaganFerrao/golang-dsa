package heap

import (
	"cmp"
	"fmt"
)

type Heap[T cmp.Ordered] struct {
	arr []T
}

func NewHeap[T cmp.Ordered]() *Heap[T] {
	return &Heap[T]{}
}

func (h *Heap[T]) heapifyUp(i int) {
	for i > 0 {
		parent := (i - 1) / 2
		if h.arr[parent] > h.arr[i] {
			h.arr[parent], h.arr[i] = h.arr[i], h.arr[parent]
			i = parent
		} else {
			break
		}
	}
}

func (h *Heap[T]) heapifyDown(i int) {
	lastIndex := len(h.arr) - 1
	for {
		leftChildIndex := 2*i + 1
		rightChildIndex := 2*i + 2

		minIndex := i

		if leftChildIndex <= lastIndex && h.arr[leftChildIndex] < h.arr[minIndex] {
			minIndex = leftChildIndex
		}

		if rightChildIndex <= lastIndex && h.arr[rightChildIndex] < h.arr[minIndex] {
			minIndex = rightChildIndex
		}

		if minIndex == i {
			break
		}

		h.arr[i], h.arr[minIndex] = h.arr[minIndex], h.arr[i]
		i = minIndex
	}
}

func (h *Heap[T]) Insert(key T) {
	h.arr = append(h.arr, key)
	h.heapifyUp(len(h.arr) - 1)
}

func (h *Heap[T]) UpdateAtIndex(i int, key T) error {
	if i > len(h.arr)-1 || i < 0 {
		return fmt.Errorf("Index out of bounds")
	}

	if h.arr[i] == key {
		return nil
	}

	h.arr[i] = key

	parentIndex := (i - 1) / 2
	leftChildIndex := 2*i + 1
	rightChildIndex := 2*i + 2

	if i > 0 && h.arr[parentIndex] > h.arr[i] {
		h.heapifyUp(i)
	} else if (leftChildIndex < len(h.arr) && h.arr[i] > h.arr[leftChildIndex]) || (rightChildIndex < len(h.arr) && h.arr[i] > h.arr[rightChildIndex]) {
		h.heapifyDown(i)
	}

	return nil
}

// Time complexity: O(logN) for heapify
func (h *Heap[T]) DeleteAtIndex(i int) error {
	if i > len(h.arr)-1 || i < 0 {
		return fmt.Errorf("Index out of bounds")
	}

	h.arr[i] = h.arr[len(h.arr)-1]
	h.arr = h.arr[:len(h.arr)-1]

	parentIndex := (i - 1) / 2
	leftChildIndex := 2*i + 1
	rightChildIndex := 2*i + 2

	if i > 0 && h.arr[parentIndex] > h.arr[i] {
		h.heapifyUp(i)
	} else if (leftChildIndex < len(h.arr) && h.arr[i] > h.arr[leftChildIndex]) || (rightChildIndex < len(h.arr) && h.arr[i] > h.arr[rightChildIndex]) {
		h.heapifyDown(i)
	}

	return nil
}

// Time complexity: O(N) to find the index and O(logN) for heapify, so O(N) in total
func (h *Heap[T]) Delete(key T) error {
	index := -1
	for i, v := range h.arr {
		if key == v {
			index = i
		}
	}

	if index == -1 {
		return fmt.Errorf("Element not found")
	}
	err := h.DeleteAtIndex(index)
	if err != nil {
		return err
	}

	return nil
}

func (h *Heap[T]) Extract() (T, error) {
	if len(h.arr) == 0 {
		var zeroValue T
		return zeroValue, fmt.Errorf("No elements found")
	}

	min := h.arr[0]
	lastIndex := len(h.arr) - 1
	lastElement := h.arr[lastIndex]

	h.arr[0] = lastElement
	h.arr = h.arr[:lastIndex]
	h.heapifyDown(0)

	return min, nil
}

func (h *Heap[T]) Peek() (T, error) {
	if len(h.arr) == 0 {
		var zeroValue T
		return zeroValue, fmt.Errorf("No elements found")
	}

	return h.arr[len(h.arr)-1], nil
}
