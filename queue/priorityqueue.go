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

func (q *priorityQueue[T]) Enqueue(item Item[T]) {
	if len(q.items) == 0 {
		q.items = append(q.items, item)
		return
	}

	index := len(q.items)
	for i, v := range q.items {
		if v.Priority < item.Priority {
			index = i
			break
		}
	}

	newSlice := make([]Item[T], len(q.items)+1)
	copy(newSlice, q.items[:index])

	newSlice[index] = item
	copy(newSlice[index+1:], q.items[index:])

	q.items = newSlice
}

func (q *priorityQueue[T]) Dequeue() (*Item[T], error) {
	if len(q.items) == 0 {
		return nil, fmt.Errorf("Queue is empty")
	}

	element := q.items[0]

	newSlice := make([]Item[T], len(q.items)-1)
	copy(newSlice, q.items[1:])
	q.items = newSlice

	return &element, nil
}

func (q *priorityQueue[T]) Length() int {
	return len(q.items)
}

func (q *priorityQueue[T]) Clear() {
	newSlice := make([]Item[T], 0)
	q.items = newSlice
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

func (q *priorityQueue[T]) AsSlice() []Item[T] {
	newSlice := make([]Item[T], len(q.items))
	copy(newSlice, q.items)

	return q.items
}
