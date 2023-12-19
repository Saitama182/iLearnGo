package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float32 // Area() - аналогично
}

/* композиция интерфейсов */

/* type Shape interface {
	ShapeWithArea
	ShapeWithPerimetr
}


type ShapeWithArea interface {
	Area() float32
}

type ShapeWithPerimetr interface {
	Perimetr() float32
} */

/* func (s Square) Perimetr() float32 {
	return s.sideLenght * 4
} */

/* композиция интерфейсов */

type Square struct {
	sideLenght float32
}

func (s Square) Area() float32 { // Area() - аналогично
	return s.sideLenght * s.sideLenght
}

type Circle struct {
	radius float32
}

func (c Circle) Area() float32 { // Area() - аналогично
	return c.radius * c.radius * math.Pi
}

func main() {
	squere1 := Square{5}
	circle := Circle{6}

	printShapeArea(squere1)
	printShapeArea(circle)

	printInterface(squere1)
	printInterface1(circle)
	printInterface("можем сюда передать все что угодно ")
	privedenieTipov("vvv")
	privedenieTipov(22)
}

func printShapeArea(shape Shape) { //  вместо (circle Shape) указываем  (shape Shape) и так как Area() то подходит и туда и туда
	fmt.Println(shape.Area())
}

/* пустой интерфейс */

func printInterface(i interface{}) {
	fmt.Printf("%+v\n", i)
}

/* Type switch  */
func printInterface1(q interface{}) {

	switch value := q.(type) {
	case int:
		fmt.Println("int", value)
	case bool:
		fmt.Println("bool", value)
	case string:
		fmt.Println("string", value)
	default:
		fmt.Println("unknown type", value)

	}

}

/* приведение типов */

func privedenieTipov(x interface{}) {
	str, ok := x.(string)
	if !ok {
		fmt.Println("interface not string")
		return
	}
	fmt.Println(len(str))
}
