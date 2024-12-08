package homework

//Напишите программу, реализующую подобие исключений в golang. golang не поддерживает нативно исключения.
// Обработайте эти “исключения”. Реализуйте цикл от 1 до N, где N достаточно большое число и сравните время
// выполнения обработки panic и обработки ошибки. Простой пример - деление на ноль может паниковать, а может
// быть корректно обработано (вами) с возвратом ошибки.

import (
	"errors"
	"fmt"
	"time"
)

// Функция деления с возвратом ошибки
func safeDivide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("деление на ноль")
	}

	return a / b, nil
}

// Функция деления с использованием паники
func unsafeDivide(a, b int) int {
	if b == 0 {
		panic("деление на ноль")
	}

	return a / b
}

// Время выполнения с обработкой ошибки: 66.4839ms
// Время выполнения с обработкой паники: 214.7656ms
func Run562() {
	N := 1000000

	// Обработка ошибки
	start := time.Now()
	for i := 1; i <= N; i++ {
		_, err := safeDivide(10, 0)
		if err != nil {
			// fmt.Println("Обнаружена ошибка:", err)
		}
	}
	fmt.Printf("Время выполнения с обработкой ошибки: %v\n", time.Since(start))

	// Обработка паники
	start = time.Now()
	for i := 1; i <= N; i++ {
		func() {
			defer func() {
				if r := recover(); r != nil {
					// log.Println("Паника перехвачена:", r)
				}
			}()
			_ = unsafeDivide(10, 0)
		}()
	}
	fmt.Printf("Время выполнения с обработкой паники: %v\n", time.Since(start))
}
