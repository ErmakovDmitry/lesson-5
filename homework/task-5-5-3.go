package homework

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// Функция, которая читает строки из io.Reader и возвращает их в обратном порядке
func reverseReader(reader io.Reader) (result []string, err error) {
	scanner := bufio.NewScanner(reader)
	var lines []string

	// Читаем строки из reader и сохраняем их в срезе lines
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Проверяем наличие ошибок при чтении
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	// Заполняем срез result в обратном порядке
	for i := len(lines) - 1; i >= 0; i-- {
		result = append(result, lines[i])
	}

	return result, nil
}

func Run553() {
	// Строка-источник, которую будем читать
	data := "Line 1\nLine 2\nLine 3\nLine 4"

	// Создаем io.Reader из строки
	reader := strings.NewReader(data)

	// Вызываем функцию reverseReader
	reversedLines, err := reverseReader(reader)
	if err != nil {
		fmt.Printf("Ошибка при чтении: %v\n", err)
		return
	}

	// Вывод строк в обратном порядке
	for _, line := range reversedLines {
		fmt.Println(line)
	}
}
