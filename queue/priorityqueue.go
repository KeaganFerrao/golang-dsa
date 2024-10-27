package queue

import "fmt"

type Item[T any] struct {
	Item     T
	Priority int
}

type priorityQueue[T any] struct {
	items []Item[T]
}

func NewPriorityQueue[T any]() *priorityQueue[T] {
	return &priorityQueue[T]{}
}

func heapifyUp[T any](a []Item[T], i int) {
	for i > 0 {
		parent := (i - 1) / 2

		if a[parent].Priority > a[i].Priority {
			a[parent], a[i] = a[i], a[parent]
			i = parent
		} else {
			break
		}
	}
}

func heapifyDown[T any](a []Item[T], i int) {
	lastIndex := len(a) - 1
	for {
		leftChildIndex := 2*i + 1
		rightChildIndex := 2*i + 2

		index := i

		if leftChildIndex <= lastIndex && a[leftChildIndex].Priority < a[index].Priority {
			index = leftChildIndex
		}

		if rightChildIndex <= lastIndex && a[rightChildIndex].Priority < a[index].Priority {
			index = rightChildIndex
		}

		if index == i {
			break
		}

		a[i], a[index] = a[index], a[i]
		i = index
	}
}

// Time complexity: O(logN)
// In an array based implementation the time complexity would be O(N) to place the element in the
// right index
func (q *priorityQueue[T]) Enqueue(item Item[T]) {
	q.items = append(q.items, item)
	heapifyUp(q.items, len(q.items)-1)
}

// Time complexity: O(logN)
// In an array based implementation the time complexity would be O(1) since the 0th index
// would just be removed and no heapify operation would be needed
func (q *priorityQueue[T]) Dequeue() (*Item[T], error) {
	if len(q.items) == 0 {
		return nil, fmt.Errorf("Queue is empty")
	}

	element := q.items[0]

	q.items[0] = q.items[len(q.items)-1]
	q.items = q.items[:len(q.items)-1]

	heapifyDown(q.items, 0)

	return &element, nil
}

func (q *priorityQueue[T]) Length() int {
	return len(q.items)
}

func (q *priorityQueue[T]) Clear() {
	q.items = make([]Item[T], 0)
}

func (q *priorityQueue[T]) Front() (*Item[T], error) {
	if len(q.items) == 0 {
		return nil, fmt.Errorf("Queue is empty")
	}

	return &q.items[0], nil
}

func (q *priorityQueue[T]) PrintQueue() {
	for i, v := range q.items {
		fmt.Printf("Item %v: %v\n", i, v)
	}
}
