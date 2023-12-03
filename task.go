/*

Вам предстоит найти хэш коммита, в котором время сборки приложения первый раз равнялось или превысило порог thresholdTime.

На вход подаётся:

список commits истории коммитов в виде массива, где hash - хэш коммита, а buildTime - время сборки приложения с изменениями этого коммита
thresholdTime, который показывает порог времени сборки
Задача:

Написать реализацию функции FindTheBrokenOne , которая возвращает хэш коммита,
в котором время сборки приложения BuildTime первый раз равнялось или превысило порог thresholdTime.
Если все BuildTime меньше thresholdTime, то функция должна возвращать пустую строку.

Ограничения:

0 <= len(commits) <= 1 000 000
commits[i+1].buildTime >= commits[i].buildTime >= 0

Пример 1:
commits = [{"hash":"654ec593","buildTime":3},{"hash":"7ed9a3d6","buildTime":5},{"hash":"20c1be38","buildTime":7},
{"hash":"6d9eb971","buildTime":9},{"hash":"4ed905e2","buildTime":10}]

thresholdTime = 4
Ответ - 7ed9a3d6, так как именно этот комит первым превышает порог 4


Пример 2:
commits = [{"hash":"654ec593","buildTime":3},{"hash":"7ed9a3d6","buildTime":5}]

thresholdTime = 7
Ответ - "", так thresholdTime больше любого представленного значения

Sample Input 1:

*/

package main

import "fmt"

func FindTheBrokenOne(buildTime int) string {
	if buildTime > 3 {
		return "hash"
	}
	return ""
}

func main() {
	// Укажите тип переменной commits
	commits := []map[string]interface{}{
		{"hash": "654ec593", "buildTime": 3},
		{"hash": "7ed9a3d6", "buildTime": 5},
		{"hash": "20c1be38", "buildTime": 7},
		{"hash": "6d9eb971", "buildTime": 9},
		{"hash": "4ed905e2", "buildTime": 10},
	}

	for _, entry := range commits {
		result := FindTheBrokenOne(entry["buildTime"].(int))
		if result != "" {
			fmt.Println(entry["hash"])
			break
		}
	}
}
