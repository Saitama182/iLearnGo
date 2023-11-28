package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	f1 := 1
	f2 := 1
	i := 3

	fmt.Print(f1, " ", f2, " ")

	for ; f2 < n; i++ {
		f1, f2 = f2, f1+f2
		fmt.Print(f2, " ")
	}
}
