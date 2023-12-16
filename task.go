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

type Commit struct {
	Hash      string
	BuildTime int
}

func FindTheBrokenOne(commits []Commit, thresholdTime int) string {
	for _, commit := range commits {
		if commit.BuildTime >= thresholdTime {
			return commit.Hash
		}
	}
	return ""
}

func quicksort(commits []Commit) []Commit {

	if len(commits) <= 1 {
		return commits
	}

	pivot := commits[len(commits)/2]

	var left, right, equal []Commit

	for _, value := range commits {

		switch {
		case value.BuildTime < pivot.BuildTime:
			left = append(left, value)
		case value.BuildTime == pivot.BuildTime:
			equal = append(equal, value)
		case value.BuildTime > pivot.BuildTime:
			right = append(right, value)
		}

	}

	return append(append(quicksort(left), equal...), quicksort(right)...)

}

func binarySearchIsSorted(commits []Commit) bool {
	low, high := 0, len(commits)-1

	for low < high {
		mid := (low + high) / 2
		if commits[mid].BuildTime > commits[mid+1].BuildTime {
			return false
		}
		if commits[mid].BuildTime < commits[mid-1].BuildTime {
			return false
		}

		if commits[mid].BuildTime > commits[low].BuildTime {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return true
}

func main() {
	commits := []Commit{
		{Hash: "654ec593", BuildTime: 112},
		{Hash: "7ed9a3d6", BuildTime: 5},
		{Hash: "20c1be38", BuildTime: 7},
		{Hash: "6d9eb971", BuildTime: 9},
		{Hash: "4ed905e2", BuildTime: 10},
	}

	if len(commits) < 0 || len(commits) > 1000000 {
		fmt.Println("Некорректная длина среза commits.")
		return
	}

	sortCom := quicksort(commits)

	if sortCom[0].BuildTime < 1 {
		fmt.Println("Неудовлетворяет условию commits[i+1].buildTime >= commits[i].buildTime >= 0")
		return
	}

	if !binarySearchIsSorted(commits) {
		fmt.Println("Неудовлетворяет условию commits[i+1].buildTime >= commits[i].buildTime >= 0")

		return
	}

	thresholdTime := 4

	res := FindTheBrokenOne(commits, thresholdTime)

	fmt.Println(res)
}
