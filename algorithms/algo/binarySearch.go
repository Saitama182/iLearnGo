package main

import (
	"fmt"
	"os"
)

func binSearch(key int, arr []int) bool {
	lo := 0
	hi := len(arr) - 1

	for lo <= hi {
		mid := lo + (hi-lo)/2
		if key < arr[mid] {
			hi = mid - 1
		} else if key > arr[mid] {
			lo = mid + 1
		} else {
			return true
		}
	}
	return false
}

func main() {
	// Путь к вашему текстовому файлу
	filePath := "testARR.txt"

	// Чтение данных из файла
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("ОШИБКА - ", err)
		os.Exit(1)
	}

	intArray := splitNumbers(string(fileData))

	var x int
	fmt.Println("Введите число:")

	fmt.Scan(&x)

	// Поиск числа в массиве
	if binSearch(x, intArray) {
		fmt.Printf("Число %d найдено в файле.\n", x)
	} else {
		fmt.Printf("Число %d  не найдено в файле.\n", x)
	}
}

func splitNumbers(input string) []int {
	var result []int
	currentNumber := 0
	for _, char := range input {
		if char >= '0' && char <= '9' {
			// Если символ - цифра, добавляем его к текущему числу
			currentNumber = currentNumber*10 + int(char-'0')
		} else {
			// Если символ не цифра, сохраняем текущее число и сбрасываем его
			if currentNumber != 0 {
				result = append(result, currentNumber)
				currentNumber = 0
			}
		}
	}
	// Проверяем последнее число после завершающего символа
	if currentNumber != 0 {
		result = append(result, currentNumber)
	}
	return result
}
