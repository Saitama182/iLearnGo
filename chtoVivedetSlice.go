package main

import "fmt"

func main() {
	var (
		foo []int
		bar []int
	)
	foo = append(foo, 1)
	fmt.Printf("Слайс1 %d  длинна слайса %d  обьем слайса %d \n", foo, len(foo), cap(foo))
	foo = append(foo, 2)
	fmt.Printf("Слайс1 %d  длинна слайса %d  обьем слайса %d \n", foo, len(foo), cap(foo))
	foo = append(foo, 3)
	fmt.Printf("Слайс1 %d  длинна слайса %d  обьем слайса %d \n", foo, len(foo), cap(foo))
	bar = append(foo, 4)
	fmt.Printf("Слайс1 %d  длинна слайса %d  обьем слайса %d \n", foo, len(foo), cap(foo))
	fmt.Printf("Слайс2 %d  длинна слайса %d  обьем слайса %d \n", bar, len(bar), cap(bar))
	foo = append(foo, 5)
	fmt.Printf("Слайс1 %d  длинна слайса %d  обьем слайса %d \n", foo, len(foo), cap(foo))
	fmt.Println(foo, bar)
}
