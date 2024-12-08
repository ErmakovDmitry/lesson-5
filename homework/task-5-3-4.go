package homework

import (
	"fmt"
	"io"
	"os"
)

func copyFile(src, dst string) error {
	// Открытие исходного файла
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	// Создание нового файла назначения
	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	// Копирование содержимого
	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return err
	}

	// Необходимый шаг для безопасного завершения записи в файл
	err = destFile.Sync()
	if err != nil {
		return err
	}

	return nil
}

func Run534() {
	sourcePath := "c:/tmp/source.txt"
	destinationPath := "c:/tmp/destination.txt"

	err := copyFile(sourcePath, destinationPath)
	if err != nil {
		fmt.Printf("Ошибка копирования файла: %v\n", err)
	} else {
		fmt.Println("Файл успешно скопирован.")
	}
}
