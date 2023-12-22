package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// FileReader - тип, который реализует интерфейс io.Reader
type FileReader struct {
	file *os.File
}

// NewFileReader - конструктор для FileReader
func NewFileReader(filename string) (*FileReader, error) {
	// Открываем файл по указанному имени
	file, err := os.Open(filename)
	if err != nil {
		return nil, err // Возвращаем ошибку, если не удалось открыть файл
	}

	// Возвращаем экземпляр FileReader с открытым файлом
	return &FileReader{file: file}, nil
}

// Read - реализация метода интерфейса io.Reader
func (fr *FileReader) Read(p []byte) (n int, err error) {
	// Читаем данные из файла в переданный буфер
	return fr.file.Read(p)
}

// Close - метод для закрытия файла
func (fr *FileReader) Close() error {
	// Закрываем файл
	return fr.file.Close()
}

func main() {
	filename := "test.txt"
	var count int

	// Создаем экземпляр FileReader с указанием имени файла
	fileReader, err := NewFileReader(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer fileReader.Close() // Закрываем файл при выходе из функции main (defer гарантирует, что это произойдет)

	// Создаем буфер для чтения данных из файла
	buffer := make([]byte, 1024)
	for {
		// Читаем данные из файла в буфер
		n, err := fileReader.Read(buffer)
		if err == io.EOF {
			break // Если достигнут конец файла, завершаем цикл
		}
		if err != nil {
			fmt.Println("Error reading from file:", err)
			break // Если произошла ошибка чтения, завершаем цикл
		}

		// Преобразуем буфер в строку и разбиваем её на числа
		numbers := strings.Split(string(buffer[:n]), ";")

		// Ищем число 0 среди чисел
		for _, numStr := range numbers {
			count++
			num, err := strconv.Atoi(numStr)
			if err != nil {
				// Обработка ошибки преобразования строки в число
				continue
			}

			if num == 0 {
				fmt.Println(count)
				return
			}
		}
	}
}
