package homework

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

// Структура данных, которую мы ожидаем в JSON
type Data struct {
	X int    `json:"x"`
	Y string `json:"y"`
	Z bool   `json:"z"`
}

func Run543() {
	filename := "homework/task-5-4-3_data.json"

	// Открытие файла
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Ошибка при открытии файла: %v", err)
	}
	defer file.Close()

	// Чтение содержимого файла
	bytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Ошибка при чтении содержимого файла: %v", err)
	}

	// Анмаршалинг JSON в структуру
	var data Data
	if err := json.Unmarshal(bytes, &data); err != nil {
		log.Fatalf("Ошибка анмаршалинга JSON: %v", err)
	}

	// Печать успешно анмаршалированных данных
	fmt.Printf("Успешно прочитан JSON из файла %s и преобразован в структуру:\n", filename)
	fmt.Printf("X: %d, Y: %s, Z: %t\n", data.X, data.Y, data.Z)
}
