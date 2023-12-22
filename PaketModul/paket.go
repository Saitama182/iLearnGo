package main

import (
	"fmt"
	"go-ninja/basic/shape"

	"github.com/zhashkevych/scheduler"
)

func main() {

	fmt.Println("Привет")

	square := shape.NewSquare(4)
	fmt.Println(square)

	sc := scheduler.NewScheduler()
	fmt.Println(sc)
}
