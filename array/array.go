package array

import "cmp"

func LinearSearch[T cmp.Ordered](a []T, key T) int {
	for i, v := range a {
		if v == key {
			return i
		}
	}

	return -1
}

func BinarySearch[T cmp.Ordered](a []T, key T) int {
	start := 0
	end := len(a) - 1

	for start <= end {
		middle := (start + end) / 2
		if key == a[middle] {
			return middle
		}

		if key > a[middle] {
			start = middle + 1
		} else {
			end = middle - 1
		}
	}

	return -1
}

func BinarySearchRecursive[T cmp.Ordered](a []T, key T, start int, end int) int {
	for start <= end {
		middle := (start + end) / 2
		if key == a[middle] {
			return middle
		}

		if key > a[middle] {
			return BinarySearchRecursive(a, key, middle+1, end)
		} else {
			return BinarySearchRecursive(a, key, start, middle-1)
		}
	}

	return -1
}

// 2 for loops, compare i and i - 1 elements and swap if i less than i - 1
func InsertionSort[T cmp.Ordered](a []T) {
	for i := range a {
		for j := i; j > 0; j-- {
			if a[j] < a[j-1] {
				temp := a[j]
				a[j] = a[j-1]
				a[j-1] = temp
			} else {
				break
			}
		}
	}
}

// Loop through the whole array and find the smallest element, then swap it
// with the first element in the unsorted part of the array, continue the process
func SelectionSort[T cmp.Ordered](a []T) {
	for i := 0; i < len(a); i++ {
		smallestIndex := i
		toSwapp := false
		for j := i; j < len(a); j++ {
			if a[j] < a[smallestIndex] {
				smallestIndex = j
				toSwapp = true
			}
		}
		if toSwapp {
			temp := a[i]
			a[i] = a[smallestIndex]
			a[smallestIndex] = temp
		}
	}
}

// Compares adjacent elements and takes the largest element foward(Bubbles it up)
// Can have more swaps as compared to selection sort
func BubbleSort[T cmp.Ordered](a []T) {
	for i := len(a) - 1; i >= 0; i-- {
		swapped := false
		for j := 0; j < i; j++ {
			if a[j] > a[j+1] {
				temp := a[j]
				a[j] = a[j+1]
				a[j+1] = temp
				swapped = true
			}
		}
		if !swapped {
			return
		}
	}
}

func merge[T cmp.Ordered](a []T, left int, mid int, right int) {
	n1 := mid - left + 1 //length of left side
	n2 := right - mid    //length of right side

	//create arrays of specific lengths to hold the temp data for merging
	l := make([]T, n1)
	r := make([]T, n2)

	//copy the left data into temp1
	for i := 0; i < n1; i++ {
		l[i] = a[left+i]
	}

	//copy the right data into temp2
	for j := 0; j < n2; j++ {
		r[j] = a[mid+1+j]
	}

	//merge the two arrays into 1 sorted array
	var i, j int // i, j points to the first element in the 2 temp arrays
	k := left    // k points element to the main array reference

	// move the 2 pointers i and j and add elements to the main array in sorted order
	for i < n1 && j < n2 {
		if l[i] <= r[j] {
			a[k] = l[i]
			i++
		} else {
			a[k] = r[j]
			j++
		}
		k++
	}

	//copy any remaining elements from the temp arrays to the main array
	for i < n1 {
		a[k] = l[i]
		i++
		k++
	}
	for j < n2 {
		a[k] = r[j]
		j++
		k++
	}
}

func MergeSort[T cmp.Ordered](a []T, left int, right int) {
	if left >= right {
		return
	}

	// split the arrays until only 1 element is left
	mid := (left + right) / 2
	MergeSort(a, left, mid)
	MergeSort(a, mid+1, right)

	merge(a, left, mid, right)
}

func partition[T cmp.Ordered](a []T, low int, high int) int {
	pivot := a[high]

	i := low - 1
	for j := low; j < high; j++ {
		if a[j] < pivot {
			i++
			temp := a[i]
			a[i] = a[j]
			a[j] = temp
		}
	}

	temp := a[i+1]
	a[i+1] = a[high]
	a[high] = temp

	return i + 1
}

func QuickSort[T cmp.Ordered](a []T, low int, high int) {
	if low < high {
		pi := partition(a, low, high)

		QuickSort(a, low, pi-1)
		QuickSort(a, pi+1, high)
	}
}

// Build a max heap
// Time complexity: O(logN), since a heap represents a complete binary tree
// so the heaight is always maintained at O(logN)
func heapifyDown[T cmp.Ordered](arr []T, i int, lastIndex int) {
	for {
		leftChildIndex := 2*i + 1
		rightChildIndex := 2*i + 2

		maxIndex := i

		if leftChildIndex <= lastIndex && arr[leftChildIndex] > arr[maxIndex] {
			maxIndex = leftChildIndex
		}

		if rightChildIndex <= lastIndex && arr[rightChildIndex] > arr[maxIndex] {
			maxIndex = rightChildIndex
		}

		if maxIndex == i {
			break
		}

		arr[i], arr[maxIndex] = arr[maxIndex], arr[i]
		i = maxIndex
	}
}

// Time complexity: O(NlogN)
// Space complexity: O(1)
// O(logN) for heapifying and O(N) for looping throught the array
// In place sorting, Not stable(Does not preserve initial order)
func HeapSort[T cmp.Ordered](a []T) {
	// Build the max heap, we start from n/2-1 since all the elements after
	// this index are all leaf nodes of the tree and they are by default
	// heapified, they do not require heapifyingDown
	for i := len(a)/2 - 1; i >= 0; i-- {
		heapifyDown(a, i, len(a)-1)
	}

	// Replace the root(max) node with the last element
	// Heapify the unsorted part of the tree
	// Do this untill the whole array is sorted
	for i := len(a) - 1; i > 0; i-- {
		a[0], a[i] = a[i], a[0]
		heapifyDown(a, 0, i-1)
	}
}
