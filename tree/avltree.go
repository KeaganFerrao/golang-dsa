package tree

import "cmp"

type AvlNode[T cmp.Ordered] struct {
	Data   T
	Left   *AvlNode[T]
	Right  *AvlNode[T]
	Height int
}

type AvlTree[T cmp.Ordered] struct {
	Root *AvlNode[T]
}

func NewAvlTree[T cmp.Ordered]() *AvlTree[T] {
	return &AvlTree[T]{}
}

func (avl *AvlTree[T]) PreOrderTraversal(node *AvlNode[T], cb func(*AvlNode[T])) {
	if node == nil {
		return
	}

	cb(node)
	avl.PreOrderTraversal(node.Left, cb)
	avl.PreOrderTraversal(node.Right, cb)
}

func height[T cmp.Ordered](node *AvlNode[T]) int {
	if node == nil {
		return 0
	}

	return node.Height
}

func getBalanceFactor[T cmp.Ordered](node *AvlNode[T]) int {
	if node == nil {
		return 0
	}

	return height(node.Left) - height(node.Right)
}

func rightRotate[T cmp.Ordered](y *AvlNode[T]) *AvlNode[T] {
	x := y.Left
	t2 := x.Right

	x.Right = y
	y.Left = t2

	y.Height = 1 + max(height(y.Left), height(y.Right))
	x.Height = 1 + max(height(x.Left), height(x.Right))

	return x
}

func leftRotate[T cmp.Ordered](x *AvlNode[T]) *AvlNode[T] {
	y := x.Right
	t2 := y.Left

	y.Left = x
	x.Right = t2

	x.Height = 1 + max(height(x.Left), height(x.Right))
	y.Height = 1 + max(height(y.Left), height(y.Right))

	return y
}

func (avl *AvlTree[T]) Insert(node *AvlNode[T], key T) *AvlNode[T] {
	// Perform normal bst insertion
	if node == nil {
		return &AvlNode[T]{Data: key, Height: 1}
	}

	if key < node.Data {
		node.Left = avl.Insert(node.Left, key)
	} else if key > node.Data {
		node.Right = avl.Insert(node.Right, key)
	} else {
		return node
	}

	// Update the height
	node.Height = 1 + max(height(node.Left), height(node.Right))

	// After normal bst insertion check the balance factor
	// If the balance factor is not in (-1, 0, 1) we need to perform rotations
	balance := getBalanceFactor(node)

	// Left Left case
	if balance > 1 && key < node.Left.Data {
		return rightRotate(node)
	}

	// Right Right case
	if balance < -1 && key > node.Right.Data {
		return leftRotate(node)
	}

	// Left Right case
	if balance > 1 && key > node.Left.Data {
		node.Left = leftRotate(node.Left)
		return rightRotate(node)
	}

	// Right Left case
	if balance < -1 && key < node.Right.Data {
		node.Right = rightRotate(node.Right)
		return leftRotate(node)
	}

	// Return unchanged pointer
	return node
}

func getPreOrderSuccessor[T cmp.Ordered](node *AvlNode[T]) *AvlNode[T] {
	temp := node.Right
	for temp != nil && temp.Left != nil {
		temp = temp.Left
	}

	return temp
}

func (avl *AvlTree[T]) Delete(node *AvlNode[T], key T) *AvlNode[T] {
	if node == nil {
		return node
	}

	if key < node.Data {
		node.Left = avl.Delete(node.Left, key)
	} else if key > node.Data {
		node.Right = avl.Delete(node.Right, key)
	} else {
		if node.Left == nil {
			return node.Right
		}

		if node.Right == nil {
			return node.Left
		}

		successor := getPreOrderSuccessor(node)
		node.Data = successor.Data
		node.Right = avl.Delete(node.Right, successor.Data)
	}

	if node == nil {
		return node
	}

	// Update the height
	node.Height = 1 + max(height(node.Left), height(node.Right))

	// After normal bst insertion check the balance factor
	// If the balance factor is not in (-1, 0, 1) we need to perform rotations
	balance := getBalanceFactor(node)

	// Left Left case
	if balance > 1 && getBalanceFactor(node.Left) >= 0 {
		return rightRotate(node)
	}

	// Right Right case
	if balance < -1 && getBalanceFactor(node.Right) <= 0 {
		return leftRotate(node)
	}

	// Left Right case
	if balance > 1 && getBalanceFactor(node.Left) < 0 {
		node.Left = leftRotate(node.Left)
		return rightRotate(node)
	}

	// Right Left case
	if balance < -1 && getBalanceFactor(node.Right) > 0 {
		node.Right = rightRotate(node.Right)
		return leftRotate(node)
	}

	return node
}
