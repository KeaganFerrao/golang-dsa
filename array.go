package main

func LinearSearch(a []int, key int) int {
	for i, v := range a {
		if v == key {
			return i
		}
	}

	return -1
}

func BinarySearch(a []int, key int) int {
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

func BinarySearchRecursive(a []int, key int, start int, end int) int {
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
func InsertionSort(a []int) {
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
func SelectionSort(a []int) {
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
func BubbleSort(a []int) {
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

func merge(a []int, left int, mid int, right int) {
	n1 := mid - left + 1 //length of left side
	n2 := right - mid    //length of right side

	//create arrays of specific lengths to hold the temp data for merging
	l := make([]int, n1)
	r := make([]int, n2)

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

func MergeSort(a []int, left int, right int) {
	if left >= right {
		return
	}

	// split the arrays until only 1 element is left
	mid := (left + right) / 2
	MergeSort(a, left, mid)
	MergeSort(a, mid+1, right)

	merge(a, left, mid, right)
}

func partition(a []int, low int, high int) int {
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

func QuickSort(a []int, low int, high int) {
	if low < high {
		pi := partition(a, low, high)

		QuickSort(a, low, pi-1)
		QuickSort(a, pi+1, high)
	}
}
