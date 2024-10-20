package stack

import "fmt"

type stack[T any] struct {
	items []T
}

func NewStack[T any]() *stack[T] {
	return &stack[T]{}
}

func (s *stack[T]) Length() int {
	return len(s.items)
}

func (s *stack[T]) Push(element ...T) {
	s.items = append(s.items, element...)
}

func (s *stack[T]) Pop() (*T, error) {
	if len(s.items) == 0 {
		return nil, fmt.Errorf("stack is empty")
	}

	poppedElement := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]

	return &poppedElement, nil
}

func (s *stack[T]) Top() (*T, error) {
	if len(s.items) == 0 {
		return nil, fmt.Errorf("stack is empty")
	}

	return &s.items[len(s.items)-1], nil
}

func (s *stack[T]) PrintStack() {
	for i, v := range s.items {
		fmt.Printf("Item %v: %v\n", i, v)
	}
}

func (s *stack[T]) Clear() {
	if len(s.items) == 0 {
		return
	}
	s.items = make([]T, 0)
}
