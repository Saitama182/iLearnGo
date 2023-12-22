package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Structure описывает структуру данных из файла structure-20190514T0000.json
type Structure struct {
	MinItems    int    `json:"minItems"`
	Schema      string `json:"$schema"`
	Description string `json:"description"`
	Title       string `json:"title"`
	Type        string `json:"type"`
	Items       Items  `json:"items"`
}

// Items описывает структуру элементов внутри массива
type Items struct {
	Description string                 `json:"description"`
	Type        string                 `json:"type"`
	Properties  map[string]interface{} `json:"properties"`
	Required    []string               `json:"required"`
}

func main() {
	// Чтение структуры из файла structure-20190514T0000.json
	structureData, err := ioutil.ReadFile("structure-20190514T0000.json")
	if err != nil {
		fmt.Println("Ошибка чтения файла structure-20190514T0000.json:", err)
		return
	}

	// Декодирование структуры
	var structure Structure
	err = json.Unmarshal(structureData, &structure)
	if err != nil {
		fmt.Println("Ошибка декодирования структуры JSON:", err)
		return
	}

	// Чтение данных из файла data-20190514T0100.json
	data, err := ioutil.ReadFile("data-20190514T0100.json")
	if err != nil {
		fmt.Println("Ошибка чтения файла data-20190514T0100.json:", err)
		return
	}

	// Декодирование данных
	var dataArray []map[string]interface{}
	err = json.Unmarshal(data, &dataArray)
	if err != nil {
		fmt.Println("Ошибка декодирования данных JSON:", err)
		return
	}

	// Вычисление суммы полей global_id
	sumGlobalID := 0
	for _, item := range dataArray {
		globalID, ok := item["global_id"].(float64)
		if ok {
			sumGlobalID += int(globalID)
		}
	}

	// Вывод суммы
	fmt.Println("Сумма полей global_id:", sumGlobalID)
}
