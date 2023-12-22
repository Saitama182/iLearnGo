package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Create("testARR.txt")

	if err != nil {
		fmt.Println("ОШИБКА - ", err)
		os.Exit(1)
	}

	defer file.Close()

	arr := make([]int, 0, 1000)

	for i := 0; i < 1001; i++ {
		arr = append(arr, i)
	}

	for i, num := range arr {
		// Добавляем пробел после каждого числа (кроме последнего)
		if i < len(arr)-1 {
			fmt.Fprintf(file, "%d ", num)
		} else {
			// Если это последнее число, добавляем перевод строки
			fmt.Fprintf(file, "%d\n", num)
		}
	}

	fmt.Println("Массив успешно записан в текстовый файл.")

}
