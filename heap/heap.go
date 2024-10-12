package heap

import (
	"cmp"
	"fmt"
)

type Heap[T cmp.Ordered] struct {
	Arr      []T
	HeapType string //"min" or "max"
}

func NewHeap[T cmp.Ordered](heapType string) *Heap[T] {
	return &Heap[T]{
		HeapType: heapType,
	}
}

func (h *Heap[T]) Length() int {
	return len(h.Arr)
}

func (h *Heap[T]) Clear() {
	h.Arr = make([]T, 0)
}

func (h *Heap[T]) heapifyUp(i int) {
	for i > 0 {
		parent := (i - 1) / 2
		if h.HeapType == "min" {
			if h.Arr[parent] > h.Arr[i] {
				h.Arr[parent], h.Arr[i] = h.Arr[i], h.Arr[parent]
				i = parent
			} else {
				break
			}
		} else {
			if h.Arr[parent] < h.Arr[i] {
				h.Arr[parent], h.Arr[i] = h.Arr[i], h.Arr[parent]
				i = parent
			} else {
				break
			}
		}
	}
}

func (h *Heap[T]) heapifyDown(i int) {
	lastIndex := len(h.Arr) - 1
	for {
		leftChildIndex := 2*i + 1
		rightChildIndex := 2*i + 2

		index := i

		if h.HeapType == "min" {
			if leftChildIndex <= lastIndex && h.Arr[leftChildIndex] < h.Arr[index] {
				index = leftChildIndex
			}

			if rightChildIndex <= lastIndex && h.Arr[rightChildIndex] < h.Arr[index] {
				index = rightChildIndex
			}
		} else {
			if leftChildIndex <= lastIndex && h.Arr[leftChildIndex] > h.Arr[index] {
				index = leftChildIndex
			}

			if rightChildIndex <= lastIndex && h.Arr[rightChildIndex] > h.Arr[index] {
				index = rightChildIndex
			}
		}

		if index == i {
			break
		}

		h.Arr[i], h.Arr[index] = h.Arr[index], h.Arr[i]
		i = index
	}
}

func (h *Heap[T]) Insert(key T) {
	h.Arr = append(h.Arr, key)
	h.heapifyUp(len(h.Arr) - 1)
}

func (h *Heap[T]) UpdateAtIndex(i int, key T) error {
	if i > len(h.Arr)-1 || i < 0 {
		return fmt.Errorf("Index out of bounds")
	}

	if h.Arr[i] == key {
		return nil
	}

	h.Arr[i] = key

	parentIndex := (i - 1) / 2
	leftChildIndex := 2*i + 1
	rightChildIndex := 2*i + 2

	if i > 0 && h.Arr[parentIndex] > h.Arr[i] {
		h.heapifyUp(i)
	} else if (leftChildIndex < len(h.Arr) && h.Arr[i] > h.Arr[leftChildIndex]) || (rightChildIndex < len(h.Arr) && h.Arr[i] > h.Arr[rightChildIndex]) {
		h.heapifyDown(i)
	}

	return nil
}

// Time complexity: O(logN) for heapify
func (h *Heap[T]) DeleteAtIndex(i int) error {
	if i > len(h.Arr)-1 || i < 0 {
		return fmt.Errorf("Index out of bounds")
	}

	h.Arr[i] = h.Arr[len(h.Arr)-1]
	h.Arr = h.Arr[:len(h.Arr)-1]

	parentIndex := (i - 1) / 2
	leftChildIndex := 2*i + 1
	rightChildIndex := 2*i + 2

	if i > 0 && h.Arr[parentIndex] > h.Arr[i] {
		h.heapifyUp(i)
	} else if (leftChildIndex < len(h.Arr) && h.Arr[i] > h.Arr[leftChildIndex]) || (rightChildIndex < len(h.Arr) && h.Arr[i] > h.Arr[rightChildIndex]) {
		h.heapifyDown(i)
	}

	return nil
}

// Time complexity: O(N) to find the index and O(logN) for heapify, so O(N) in total
func (h *Heap[T]) Delete(key T) error {
	index := -1
	for i, v := range h.Arr {
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
	if len(h.Arr) == 0 {
		var zeroValue T
		return zeroValue, fmt.Errorf("No elements found")
	}

	min := h.Arr[0]
	lastIndex := len(h.Arr) - 1
	lastElement := h.Arr[lastIndex]

	h.Arr[0] = lastElement
	h.Arr = h.Arr[:lastIndex]
	h.heapifyDown(0)

	return min, nil
}

func (h *Heap[T]) Peek() (T, error) {
	if len(h.Arr) == 0 {
		var zeroValue T
		return zeroValue, fmt.Errorf("No elements found")
	}

	return h.Arr[len(h.Arr)-1], nil
}
