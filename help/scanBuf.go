package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Считываем строки a и b
	a, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	b, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	// Удаляем пробелы и символ новой строки из строк a и b
	a = strings.TrimSpace(a)
	b = strings.TrimSpace(b)

	// Преобразуем строки в целые числа
	aInt, errA := strconv.Atoi(a)
	bInt, errB := strconv.Atoi(b)

	// Проверяем ошибки
	if errA != nil || errB != nil {
		fmt.Println("Ошибка преобразования строки в число")
		return
	}

	// Выводим сумму двух чисел
	fmt.Println(aInt + bInt)
}
