package linkedlist

import "fmt"

type node[T comparable] struct {
	Item T
	next *node[T]
}

type linkedList[T comparable] struct {
	head   *node[T]
	length int
}

func NewLinkedList[T comparable]() *linkedList[T] {
	return &linkedList[T]{}
}

func (ll *linkedList[T]) GetNext(node *node[T]) *node[T] {
	return node.next
}

func (ll *linkedList[T]) GetHead() *node[T] {
	return ll.head
}

func (ll *linkedList[T]) Length() int {
	return ll.length
}

func (ll *linkedList[T]) Clear() {
	ll.head = nil
	ll.length = 0
}

func (ll *linkedList[T]) Search(item T) *node[T] {
	node := ll.head
	for node != nil {
		if node.Item == item {
			return node
		}
		node = node.next
	}

	return nil
}

func (ll *linkedList[T]) Append(element T) {
	newNode := node[T]{Item: element, next: nil}

	if ll.length == 0 {
		ll.head = &newNode
		ll.length++
		return
	}

	start := ll.head
	for start.next != nil {
		start = start.next
	}
	start.next = &newNode
	ll.length++
}

func (ll *linkedList[T]) Prepend(element T) {
	newNode := node[T]{Item: element, next: nil}
	newNode.next = ll.head
	ll.head = &newNode
	ll.length++
}

func (ll *linkedList[T]) DeleteAt(index int) error {
	if index > ll.length-1 {
		return fmt.Errorf("Index out of bounds")
	}
	if index < 0 {
		return fmt.Errorf("Invalid index")
	}
	if index == 0 {
		temp := ll.head
		ll.head = ll.head.next
		temp.next = nil
		ll.length--

		return nil
	}

	nodePtr := ll.head
	idx := 0
	for nodePtr != nil {
		if index == idx+1 {
			break
		}
		nodePtr = nodePtr.next
		idx++
	}

	temp := nodePtr.next
	nodePtr.next = nodePtr.next.next
	temp.next = nil
	ll.length--

	return nil
}

func (ll *linkedList[T]) InsertAt(element T, index int) error {
	if index > ll.length {
		return fmt.Errorf("Index out of bounds")
	}

	if index < 0 {
		return fmt.Errorf("Invalid index")
	}

	if index == 0 {
		ll.Prepend(element)
		return nil
	}

	nodePtr := ll.head
	idx := 0
	for nodePtr != nil {
		if index == idx+1 {
			break
		}
		nodePtr = nodePtr.next
		idx++
	}

	newNode := node[T]{Item: element, next: nil}
	nextNode := nodePtr.next
	nodePtr.next = &newNode
	newNode.next = nextNode
	ll.length++

	return nil
}

func (ll *linkedList[T]) Reverse() {
	var prev, next *node[T]

	curr := ll.head

	for curr != nil {
		next = curr.next
		curr.next = prev
		prev = curr
		curr = next
	}

	ll.head = prev
}

func (ll *linkedList[T]) PrintList() {
	node := ll.head

	for node != nil {
		fmt.Printf("Element: %v\n", node.Item)
		node = node.next
	}
}
