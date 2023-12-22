package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Открываем файл для чтения
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer file.Close()

	// Создаем новый сканер для чтения файла построчно
	scanner := bufio.NewScanner(file)

	// Устанавливаем максимальный размер буфера для сканера, чтобы он мог обработать очень длинные строки
	const maxScanTokenSize = 64 * 1024 * 1024 // Например, установим максимальный размер строки в 64 МБ
	scanner.Buffer(make([]byte, maxScanTokenSize), maxScanTokenSize)

	// Считываем строки из файла
	for scanner.Scan() {
		str := scanner.Text()
		parts := strings.Split(str, ";")

		var count int

		// Проходим по всем элементам
		for _, part := range parts {
			// Пробуем преобразовать каждый элемент в число
			number, err := strconv.Atoi(part)
			count++
			if err == nil && number == 0 {
				fmt.Println(count)
				return
			}
		}

		fmt.Println("Число, равное 0, не найдено.")
	}

	// Проверяем наличие ошибок после завершения сканирования
	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка при сканировании файла:", err)
	}
}
