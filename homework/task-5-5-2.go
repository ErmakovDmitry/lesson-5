package homework

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Run552() {
	// Имя файла для чтения
	filename := "homework/task-5-5-2_data.txt"

	// Открытие файла
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Ошибка при открытии файла: %v", err)
	}
	defer file.Close()

	// Используем сканер для чтения файла построчно
	scanner := bufio.NewScanner(file)

	// Создаем слайс деферированных функций
	var deferFuncs []func()

	for scanner.Scan() {
		// Копируем строку в локальную переменную, чтобы избежать замыкания на переменной цикла
		line := scanner.Text()
		// Добавляем функцию в слайс с использованием `defer`
		deferFuncs = append(deferFuncs, func() {
			fmt.Println(line)
		})
	}

	// Проверка на наличие ошибок во время сканирования
	if err := scanner.Err(); err != nil {
		log.Fatalf("Ошибка при чтении файла: %v", err)
	}

	// Выполнение отложенных функций в обратном порядке
	for i := len(deferFuncs) - 1; i >= 0; i-- {
		deferFuncs[i]()
	}
}
