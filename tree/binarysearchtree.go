package tree

import (
	"cmp"
)

type binarySearchTree[T cmp.Ordered] struct {
	Root *Node[T]
}

func NewBinarySearchTree[T cmp.Ordered]() *binarySearchTree[T] {
	return &binarySearchTree[T]{}
}

func (bst *binarySearchTree[T]) InOrderTraversal(node *Node[T], cb func(*Node[T])) {
	if node == nil {
		return
	}

	bst.InOrderTraversal(node.Left, cb)
	cb(node)
	bst.InOrderTraversal(node.Right, cb)
}

func (bst *binarySearchTree[T]) PreOrderTraversal(node *Node[T], cb func(*Node[T])) {
	if node == nil {
		return
	}

	cb(node)
	bst.PreOrderTraversal(node.Left, cb)
	bst.PreOrderTraversal(node.Right, cb)
}

func (bst *binarySearchTree[T]) PostOrderTraversal(node *Node[T], cb func(*Node[T])) {
	if node == nil {
		return
	}

	bst.PostOrderTraversal(node.Left, cb)
	bst.PostOrderTraversal(node.Right, cb)
	cb(node)
}

func (bst *binarySearchTree[T]) LevelOrderTraversal(node *Node[T], cb func(*Node[T])) {
	queue := make([]*Node[T], 0)
	queue = append(queue, node)

	for len(queue) > 0 {
		temp := queue[0]
		cb(temp)
		queue = queue[1:]

		if temp.Left != nil {
			queue = append(queue, temp.Left)
		}
		if temp.Right != nil {
			queue = append(queue, temp.Right)
		}
	}
}

func (bst *binarySearchTree[T]) InsertRecursive(key T, root *Node[T]) *Node[T] {
	if root == nil {
		return &Node[T]{Data: key}
	}

	// No duplicates
	if key == root.Data {
		return root
	}

	if key < root.Data {
		root.Left = bst.InsertRecursive(key, root.Left)
	} else if key > root.Data {
		root.Right = bst.InsertRecursive(key, root.Right)
	}

	return root
}

func (bst *binarySearchTree[T]) Insert(key T) {
	node := &Node[T]{Data: key}
	if bst.Root == nil {
		bst.Root = node
		return
	}

	var parent *Node[T]
	curr := bst.Root
	for curr != nil {
		parent = curr
		if key < curr.Data {
			curr = curr.Left
		} else if key > curr.Data {
			curr = curr.Right
		} else {
			// No duplicates
			return
		}
	}

	if key < parent.Data {
		parent.Left = node
	} else if key > parent.Data {
		parent.Right = node
	}
}

func (bst *binarySearchTree[T]) Height(root *Node[T]) int {
	if root == nil {
		return 0
	}
	return 1 + max(bst.Height(root.Left), bst.Height(root.Right))
}

func (bst *binarySearchTree[T]) Size(root *Node[T]) int {
	if root == nil {
		return 0
	}

	return 1 + bst.Size(root.Left) + bst.Size(root.Right)
}

// Leftmost node of the right subtree. Which node will come next in the inorder traversal
func getInorderSuccessor[T cmp.Ordered](node *Node[T]) *Node[T] {
	temp := node.Right
	for temp != nil && temp.Left != nil {
		temp = temp.Left
	}
	return temp
}

func (bst *binarySearchTree[T]) Delete(key T, node *Node[T]) *Node[T] {
	if bst.Root == nil {
		return nil
	}

	if key < node.Data {
		node.Left = bst.Delete(key, node.Left)
	} else if key > node.Data {
		node.Right = bst.Delete(key, node.Right)
	} else {
		if node.Left == nil {
			return node.Right
		}

		if node.Right == nil {
			return node.Left
		}

		successor := getInorderSuccessor(node)
		node.Data = successor.Data
		node.Right = bst.Delete(successor.Data, node.Right)
	}

	return node
}

func sortedArrayToBalancedBst[T cmp.Ordered](arr []T, start int, end int) *Node[T] {
	if start > end {
		return nil
	}

	mid := (start + end) / 2
	node := &Node[T]{Data: arr[mid]}

	node.Left = sortedArrayToBalancedBst(arr, start, mid-1)
	node.Right = sortedArrayToBalancedBst(arr, mid+1, end)

	return node
}

func (bst *binarySearchTree[T]) Rebalance() {
	var arr []T

	bst.InOrderTraversal(bst.Root, func(n *Node[T]) {
		arr = append(arr, n.Data)
	})

	bst.Root = sortedArrayToBalancedBst(arr, 0, len(arr)-1)
}

func (bst *binarySearchTree[T]) Search(root *Node[T], key T) *Node[T] {
	if root == nil {
		return nil
	}

	if key < root.Data {
		return bst.Search(root.Left, key)
	} else if key > root.Data {
		return bst.Search(root.Right, key)
	} else {
		return root
	}
}

// Without the need for recursion stack to traverse the tree.
// This modifies the tree while traversal but restores the original structure later
// Space complexity O(1) as compared to classic inorder O(N) due to recursion stack
// This can be used in systems where space is a concern. However it is not always
// desirable in some systems to modify the tree structure
func (bst *binarySearchTree[T]) MorrisInorderTraversal(cb func(*Node[T])) {
	curr := bst.Root

	for curr != nil {
		if curr.Left == nil {
			cb(curr)
			curr = curr.Right
		} else {
			// Find the inorder predecessor.
			// Node which comes before a node in an inorder traversal
			predessor := curr.Left
			for predessor.Right != nil && predessor.Right != curr {
				predessor = predessor.Right
			}

			if predessor.Right == nil {
				predessor.Right = curr
				curr = curr.Left
			} else {
				predessor.Right = nil
				cb(curr)
				curr = curr.Right
			}
		}
	}
}
