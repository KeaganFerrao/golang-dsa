package main

import (
	binarytree "practice/tree"
)

func main() {
	bt := binarytree.NewBinarySearchTree[int]()

	// node1 := &binarytree.Node[int]{Data: 1}
	// node2 := &binarytree.Node[int]{Data: 2}
	// node3 := &binarytree.Node[int]{Data: 3}
	// node4 := &binarytree.Node[int]{Data: 4}
	// node5 := &binarytree.Node[int]{Data: 5}
	// node6 := &binarytree.Node[int]{Data: 6}
	// node7 := &binarytree.Node[int]{Data: 7}
	// node8 := &binarytree.Node[int]{Data: 8}

	error := bt.Insert(bt.Root, 50)
	if error != nil {
		panic(error)
	}

	error = bt.Insert(bt.Root, 30)
	if error != nil {
		panic(error)
	}
	error = bt.Insert(bt.Root, 70)
	if error != nil {
		panic(error)
	}

	bt.Insert(bt.Root, 40)
	bt.Insert(bt.Root, 20)
	bt.Insert(bt.Root, 10)
	bt.Insert(bt.Root, 14)
	bt.Insert(bt.Root, 15)
	bt.Insert(bt.Root, 12)
	bt.Insert(bt.Root, 9)
	bt.Insert(bt.Root, 16)
	bt.Insert(bt.Root, 35)

	// ele := bt.Search(bt.Root, 999)
	// fmt.Printf("Ele is %v", ele)
	bt.DeleteRecursive(bt.Root, 9)
	bt.InOrderTraversal(bt.Root)
}
