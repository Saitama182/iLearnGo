package main

import "fmt"

func main() {
	/* ssl2 := make([]int, 1) */

	var ssl2 []int

	fmt.Printf("Слайс2 %d длинна слайса %d  обьем слайса %d  \n", ssl2, len(ssl2), cap(ssl2))

}
