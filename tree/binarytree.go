package tree

import "fmt"

type Node[T any] struct {
	Data  T
	Left  *Node[T]
	Right *Node[T]
}

type binaryTree[T any] struct {
	Root *Node[T]
}

func NewBinaryTree[T any]() *binaryTree[T] {
	return &binaryTree[T]{}
}

func (bt *binaryTree[T]) PreOrderTraversal(node *Node[T]) {
	if node == nil {
		return
	}

	fmt.Printf("Data: %v\n", node.Data)

	bt.PreOrderTraversal(node.Left)
	bt.PreOrderTraversal(node.Right)
}

func (bt *binaryTree[T]) InOrderTraversal(node *Node[T]) {
	if node == nil {
		return
	}

	bt.InOrderTraversal(node.Left)
	fmt.Printf("Data: %v\n", node.Data)
	bt.InOrderTraversal(node.Right)
}

func (bt *binaryTree[T]) PostOrderTraversal(node *Node[T]) {
	if node == nil {
		return
	}

	bt.InOrderTraversal(node.Left)
	bt.InOrderTraversal(node.Right)

	fmt.Printf("Data: %v\n", node.Data)
}

func (bt *binaryTree[T]) Height(node *Node[T]) int {
	if node == nil {
		return 0
	}

	return 1 + max(bt.Height(node.Left), bt.Height(node.Right))
}

func (bt *binaryTree[T]) LevelOrderTraversal(node *Node[T]) {
	queue := make([]*Node[T], 0)

	queue = append(queue, node)
	for len(queue) > 0 {
		element := queue[0]
		queue = queue[1:]
		fmt.Printf("Data: %v\n", element.Data)

		if element.Left != nil {
			queue = append(queue, element.Left)
		}
		if element.Right != nil {
			queue = append(queue, element.Right)
		}
	}
}

func (bt *binaryTree[T]) ReverseLevelOrderTraversal(node *Node[T]) {
	queue := make([]*Node[T], 0)
	stack := make([]*Node[T], 0)

	queue = append(queue, node)
	for len(queue) > 0 {
		element := queue[0]
		queue = queue[1:]
		stack = append(stack, element)

		if element.Left != nil {
			queue = append(queue, element.Left)
		}
		if element.Right != nil {
			queue = append(queue, element.Right)
		}
	}

	for len(stack) > 0 {
		element := stack[len(stack)-1]

		stack = stack[:len(stack)-1]
		fmt.Printf("Data: %v\n", element.Data)
	}
}

func (bt *binaryTree[T]) Size(node *Node[T]) int {
	if node == nil {
		return 0
	}

	return 1 + bt.Size(node.Left) + bt.Size(node.Right)
}

func (bt *binaryTree[T]) InsertInLevelOrder(node *Node[T]) {
	if bt.Root == nil {
		bt.Root = node
		return
	}

	queue := make([]*Node[T], 0)
	queue = append(queue, bt.Root)

	for len(queue) > 0 {
		element := queue[0]
		queue = queue[1:]

		if element.Left != nil {
			queue = append(queue, element.Left)
		} else {
			element.Left = node
			return
		}

		if element.Right != nil {
			queue = append(queue, element.Right)
		} else {
			element.Right = node
			return
		}
	}
}
