package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	inputedTime, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}

	splitedTime := strings.Split(inputedTime, ",")

	var firstPart, secondPart string
	if len(splitedTime) >= 2 {
		firstPart = splitedTime[0]
		secondPart = splitedTime[1]
	}

	splitedTime[1] = strings.TrimRight(splitedTime[1], "\n")

	fmt.Printf("Первая часть после Split: %s\n", firstPart)
	fmt.Printf("Вторая часть после Split: %s\n", secondPart)

}
