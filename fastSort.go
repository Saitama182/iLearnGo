package main

import "fmt"

func quicksort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[len(arr)/2]
	left := make([]int, 0)
	right := make([]int, 0)
	equal := make([]int, 0)

	for _, value := range arr {
		switch {
		case value < pivot:
			left = append(left, value)
		case value == pivot:
			equal = append(equal, value)
		case value > pivot:
			right = append(right, value)
		}
	}

	return append(append(quicksort(left), equal...), quicksort(right)...)
}

func main() {
	arr := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	fmt.Println("Unsorted:", arr)

	arr = quicksort(arr)
	fmt.Println("Sorted:", arr)
}
