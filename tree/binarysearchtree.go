package tree

import (
	"cmp"
	"fmt"
)

type binarySearchTree[T cmp.Ordered] struct {
	Root *Node[T]
}

func NewBinarySearchTree[T cmp.Ordered]() *binarySearchTree[T] {
	return &binarySearchTree[T]{}
}

func (bt *binarySearchTree[T]) PreOrderTraversal(node *Node[T]) {
	if node == nil {
		return
	}

	fmt.Printf("Data: %v\n", node.Data)

	bt.PreOrderTraversal(node.Left)
	bt.PreOrderTraversal(node.Right)
}

func (bt *binarySearchTree[T]) InOrderTraversal(node *Node[T]) {
	if node == nil {
		return
	}

	bt.InOrderTraversal(node.Left)
	fmt.Printf("Data: %v\n", node.Data)
	bt.InOrderTraversal(node.Right)
}

func (bt *binarySearchTree[T]) PostOrderTraversal(node *Node[T]) {
	if node == nil {
		return
	}

	bt.InOrderTraversal(node.Left)
	bt.InOrderTraversal(node.Right)

	fmt.Printf("Data: %v\n", node.Data)
}

func (bt *binarySearchTree[T]) Height(node *Node[T]) int {
	if node == nil {
		return 0
	}

	return 1 + max(bt.Height(node.Left), bt.Height(node.Right))
}

func (bt *binarySearchTree[T]) LevelOrderTraversal(node *Node[T]) {
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

func (bt *binarySearchTree[T]) ReverseLevelOrderTraversal(node *Node[T]) {
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

func (bt *binarySearchTree[T]) Size(node *Node[T]) int {
	if node == nil {
		return 0
	}

	return 1 + bt.Size(node.Left) + bt.Size(node.Right)
}

func (bst *binarySearchTree[T]) Insert(root *Node[T], key T) error {
	node := &Node[T]{Data: key}

	if root == nil {
		bst.Root = node
		return nil
	}

	var parent *Node[T]
	curr := root

	//Loop through each node and keep a reference to the parent as well
	for curr != nil {
		parent = curr
		if curr.Data > key {
			curr = curr.Left
		} else if curr.Data < key {
			curr = curr.Right
		} else {
			return fmt.Errorf("Duplicate key %v", key)
		}
	}

	//Parent is the node where we insert
	if parent.Data > key {
		parent.Left = node
	} else {
		parent.Right = node
	}

	return nil
}

func (bst *binarySearchTree[T]) Search(root *Node[T], key T) *Node[T] {
	if root == nil || key == root.Data {
		return root
	}

	if key < root.Data {
		return bst.Search(root.Left, key)
	} else if key > root.Data {
		return bst.Search(root.Right, key)
	}

	return nil
}

// Inorder successor: The node that comes the immediate next when we follow
// an inorder traversal => Left, Root, Right
// We do an inorder traversal, because it always results in nodes in
// ascending order, so we know which node should we choose to replace the
// node to delete with
// Here we only consider right subtree because we are sure that it has a
// right since we are checking conditions in the caller code.
func getSuccessor[T cmp.Ordered](node *Node[T]) *Node[T] {
	node = node.Right
	for node != nil && node.Left != nil {
		node = node.Left
	}
	return node
}

func (bst *binarySearchTree[T]) DeleteRecursive(root *Node[T], key T) *Node[T] {
	if root == nil {
		return nil
	}

	if key < root.Data {
		root.Left = bst.DeleteRecursive(root.Left, key)
	} else if key > root.Data {
		root.Right = bst.DeleteRecursive(root.Right, key)
	} else {
		// If only single child present
		if root.Left == nil {
			return root.Right
		}
		// If only single child present
		if root.Right == nil {
			return root.Left
		}

		// If both children present, get the inorder successor
		// Replace the data in the node with the successor
		// And delete the successor in the tree, which is
		// the leftmost node in the right subtree
		successor := getSuccessor(root)
		root.Data = successor.Data
		root.Right = bst.DeleteRecursive(root.Right, successor.Data)
	}

	return root
}

// func (bst *binarySearchTree[T]) Rebalance() {

// }

func SortedArrayToBalancedBst[T cmp.Ordered](arr []T) *binarySearchTree[T] {

	return nil
}
