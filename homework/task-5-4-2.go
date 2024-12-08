package homework

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

type Example struct {
	A int    `json:"a"`
	B string `json:"B,omitempty"`
}

func jsonToStruct(s []byte) (*Example, error) {
	if len(s) == 0 {
		return nil, errors.New("входной JSON пустой")
	}

	var example Example
	err := json.Unmarshal(s, &example)
	if err != nil {
		return nil, err
	}

	return &example, nil
}

func Run542() {
	// Пример корректного JSON
	validJSON := []byte(`{"a": 100, "B": "Hello, world!"}`)

	// Вызов функции jsonToStruct с корректным JSON
	example, err := jsonToStruct(validJSON)
	if err != nil {
		log.Fatalf("Ошибка преобразования JSON в структуру: %v", err)
	}

	fmt.Println("Преобразованный объект Example с корректным JSON:")
	fmt.Printf("A: %d, B: %s\n", example.A, example.B)

	// Пример JSON с отсутствующим полем B (чтобы проверить работу `omitempty`)
	jsonWithoutB := []byte(`{"a": 200}`)

	// Вызов функции jsonToStruct с JSON без поля B
	example2, err := jsonToStruct(jsonWithoutB)
	if err != nil {
		log.Fatalf("Ошибка преобразования JSON в структуру: %v", err)
	}

	fmt.Println("\nПреобразованный объект Example с JSON без поля B:")
	fmt.Printf("A: %d, B: %s\n", example2.A, example2.B)

	// Пример некорректного JSON для проверки обработки ошибки
	invalidJSON := []byte(`{"a": "string_instead_of_int", "B": "Hello!"}`)

	fmt.Println("\ninvalidJSON ", string(invalidJSON[:]))

	// Вызов функции jsonToStruct с некорректным JSON
	_, err = jsonToStruct(invalidJSON)
	if err != nil {
		log.Fatalf("\nОжидаемая ошибка при попытке преобразовать некорректный JSON: %v\n", err)
	}
}
