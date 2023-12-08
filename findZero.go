package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var numbers []int64

	// Открываем файл для чтения
	file, err := os.Open("test2.txt")
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer file.Close() // Важно закрыть файл после использования

	// Создаем новый сканер для чтения файла построчно
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		numbersStr := strings.Split(line, ";")
		for _, numStr := range numbersStr {
			// Проверка на пустую строку
			if numStr == "" {
				continue
			}

			num, err := strconv.ParseInt(numStr, 10, 64)
			if err != nil {
				fmt.Println("Ошибка при преобразовании строки в число:", err)
				return
			}
			numbers = append(numbers, num)
		}
	}

	for i, num := range numbers {
		if num == 0 {
			fmt.Printf("Порядковый номер: %d, Значение: %d\n", i, num)
			break
		}
	}

	// Проверяем наличие ошибок после завершения сканирования
	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка при сканировании файла:", err)
	}

}
