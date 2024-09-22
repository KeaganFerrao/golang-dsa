package main

import (
	"fmt"
	binarytree "practice/tree"
)

func main() {
	bt := binarytree.NewBinarySearchTree[int]()

	bt.Root = bt.InsertRecursive(10, bt.Root)
	bt.Root = bt.InsertRecursive(11, bt.Root)
	bt.Root = bt.InsertRecursive(8, bt.Root)
	bt.Root = bt.InsertRecursive(7, bt.Root)
	bt.Delete(11, bt.Root)

	bt.PreOrderTraversal(bt.Root, func(n *binarytree.Node[int]) {
		fmt.Printf("n.Data: %v\n", n.Data)
	})

	avl := binarytree.NewAvlTree[int]()
	avl.Root = avl.Insert(avl.Root, 10)
	avl.Root = avl.Insert(avl.Root, 11)
	avl.Root = avl.Insert(avl.Root, 8)
	avl.Root = avl.Insert(avl.Root, 7)

	avl.Root = avl.Delete(avl.Root, 11)

	avl.PreOrderTraversal(avl.Root, func(n *binarytree.AvlNode[int]) {
		fmt.Printf("AVL n.Data: %v\n", n.Data)
	})

	// ele := bt.Search(bt.Root, 999)
	// fmt.Printf("Ele is %v", ele)
	// bt.DeleteRecursive(bt.Root, 9)
	// bt.InOrderTraversal(bt.Root)
	// bt.Rebalance()

	// bt.InOrderTraversal(bt.Root, func(n *binarytree.Node[int]) {
	// 	fmt.Printf("n.Data: %v\n", n.Data)
	// })

	// bt.MorrisInorderTraversal(func(n *binarytree.Node[int]) {
	// 	fmt.Printf("n.Data: %v\n", n.Data)
	// })
}
