package main

import "fmt"

func main() {
	var arr [50]int
	for i := 0; i < len(arr); i++ {
		arr[i] = i
	}
	fmt.Println(arr)

	var low, high, guess, item, count int

	fmt.Print("Enter the number to search: ")
	fmt.Scan(&item)

	high = len(arr) - 1
	var mid int

	for low <= high {
		mid = (low + high) / 2
		guess = arr[mid]

		if guess == item {
			fmt.Println("Element found at index:", mid)
			fmt.Println("Step:", count)
			return

		} else if guess > item {
			high = mid - 1
			count++
		} else {
			low = mid + 1
			count++
		}
	}

	fmt.Println("Element not found in the array.")

}
