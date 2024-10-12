package main

import (
	"fmt"
	"practice/array"
)

func main() {
	arr := []int{1, 5, 3, 6, 2, 4, 3, 7, 34, 78, 0, 35}
	array.HeapSort(arr)

	fmt.Printf("arr: %v\n", arr)
}
