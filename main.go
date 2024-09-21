package main

import (
	"fmt"
	binarytree "practice/tree"
)

func main() {
	bt := binarytree.NewBinarySearchTree[int]()

	bt.Insert(40)
	bt.Insert(20)
	bt.Insert(10)
	bt.Insert(14)
	bt.Insert(15)
	bt.Insert(12)
	bt.Insert(9)
	bt.Insert(16)
	bt.Insert(35)

	// ele := bt.Search(bt.Root, 999)
	// fmt.Printf("Ele is %v", ele)
	// bt.DeleteRecursive(bt.Root, 9)
	// bt.InOrderTraversal(bt.Root)
	// bt.Rebalance()

	// bt.InOrderTraversal(bt.Root, func(n *binarytree.Node[int]) {
	// 	fmt.Printf("n.Data: %v\n", n.Data)
	// })

	bt.MorrisInorderTraversal(func(n *binarytree.Node[int]) {
		fmt.Printf("n.Data: %v\n", n.Data)
	})
}
