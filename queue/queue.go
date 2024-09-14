package queue

import "fmt"

type queue[T comparable] struct {
	items []T
}

func NewQueue[T comparable]() *queue[T] {
	return &queue[T]{}
}

func (q *queue[T]) Enqueue(element T) {
	q.items = append(q.items, element)
}

func (q *queue[T]) Dequeue() (*T, error) {
	if len(q.items) == 0 {
		return nil, fmt.Errorf("queue is empty")
	}

	element := q.items[0]

	newSlice := make([]T, len(q.items)-1)
	copy(newSlice, q.items[:len(q.items)-1])
	q.items = newSlice

	return &element, nil
}

func (q *queue[T]) Length() int {
	return len(q.items)
}

func (q *queue[T]) Clear() {
	if len(q.items) == 0 {
		return
	}
	q.items = make([]T, 0)
}

func (q *queue[T]) Search(key T) *T {
	for _, v := range q.items {
		if v == key {
			return &v
		}
	}
	return nil
}

func (q *queue[T]) Front() (*T, error) {
	if len(q.items) == 0 {
		return nil, fmt.Errorf("queue is empty")
	}
	return &q.items[0], nil
}

func (q *queue[T]) PrintQueue() {
	for i, v := range q.items {
		fmt.Printf("Item %v: %v\n", i, v)
	}
}

func (q *queue[T]) AsSlice() []T {
	newSlice := make([]T, len(q.items))
	copy(newSlice, q.items)

	return q.items
}
