package linkedlist

import "fmt"

type doublyNode[T comparable] struct {
	Item T
	next *doublyNode[T]
	prev *doublyNode[T]
}

type doubleLinkedList[T comparable] struct {
	head   *doublyNode[T] //8 bytes
	tail   *doublyNode[T] //8 bytes
	length int            //8 bytes
}

func NewDoublyLinkedList[T comparable]() *doubleLinkedList[T] {
	return &doubleLinkedList[T]{}
}

func (dll *doubleLinkedList[T]) GetHead() *doublyNode[T] {
	return dll.head
}

func (dll *doubleLinkedList[T]) GetTail() *doublyNode[T] {
	return dll.tail
}

func (dll *doubleLinkedList[T]) Append(element T) {
	newNode := &doublyNode[T]{Item: element}

	if dll.length == 0 {
		dll.head = newNode
		dll.tail = newNode
		dll.length++
		return
	}

	dll.tail.next = newNode
	newNode.prev = dll.tail
	dll.tail = newNode

	dll.length++
}

func (dll *doubleLinkedList[T]) Prepend(element T) {
	newNode := &doublyNode[T]{Item: element}

	if dll.length == 0 {
		dll.head = newNode
		dll.tail = newNode
		dll.length++
		return
	}

	newNode.next = dll.head
	dll.head.prev = newNode
	dll.head = newNode
	dll.length++
}

func (dll *doubleLinkedList[T]) PrintList() {
	curr := dll.head
	for curr != nil {
		fmt.Printf("Item: %v\n", curr.Item)
		curr = curr.next
	}
}

func (dll *doubleLinkedList[T]) InsertAt(element T, index int) error {
	if index < 0 || index > dll.length {
		return fmt.Errorf("Index %v out of bounds", index)
	}

	if index == 0 {
		dll.Prepend(element)
		return nil
	}

	if index == dll.length {
		dll.Append(element)
		return nil
	}

	var curr *doublyNode[T]
	idx := 0

	// Travere the linked from head or tail based on the closest index
	mid := dll.length / 2
	if index <= mid {
		curr = dll.head
		for curr != nil {
			if index == idx {
				break
			}
			curr = curr.next
			idx++
		}
	} else {
		curr = dll.tail
		idx = dll.length - 1
		for curr != nil {
			if index == idx {
				break
			}
			curr = curr.prev
			idx--
		}
	}

	newNode := &doublyNode[T]{Item: element}

	curr.prev.next = newNode
	newNode.next = curr.next
	newNode.prev = curr.prev
	curr.next.prev = newNode
	curr.next = nil
	curr.prev = nil

	dll.length++

	return nil
}

func (dll *doubleLinkedList[T]) DeleteAt(index int) error {
	if index < 0 || index >= dll.length {
		return fmt.Errorf("Index %v out of bounds", index)
	}

	if index == 0 {
		curr := dll.head
		dll.head = dll.head.next
		dll.head.prev = nil
		curr.next = nil
		curr.prev = nil

		dll.length--
		return nil
	}

	if index == dll.length-1 {
		curr := dll.tail
		dll.tail = dll.tail.prev
		dll.tail.next = nil
		curr.next = nil
		curr.prev = nil

		dll.length--
		return nil
	}

	var curr *doublyNode[T]
	idx := 0

	// Travere the linked from head or tail based on the closest index
	mid := dll.length / 2
	if index <= mid {
		curr = dll.head
		for curr != nil {
			if index == idx {
				break
			}
			curr = curr.next
			idx++
		}
	} else {
		curr = dll.tail
		idx = dll.length - 1
		for curr != nil {
			if index == idx {
				break
			}
			curr = curr.prev
			idx--
		}
	}

	curr.prev.next = curr.next
	curr.next.prev = curr.prev
	curr.next = nil
	curr.prev = nil

	dll.length--

	return nil
}

func (dll *doubleLinkedList[T]) Reverse() {
	curr := dll.head
	dll.tail = curr

	var prev, next *doublyNode[T]
	for curr != nil {
		next = curr.next
		curr.next = prev
		curr.prev = next

		prev = curr
		curr = next
	}

	dll.head = prev
}
