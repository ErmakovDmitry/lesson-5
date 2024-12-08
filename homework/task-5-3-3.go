package homework

import (
	"fmt"
	"io"
)

type MyReader struct {
	data []byte
	pos  int
}

type MyWriter struct {
	data []byte
	pos  int
}

func NewReaderFromBuffer(buffer []byte) *MyReader {
	return &MyReader{
		data: buffer,
		pos:  0,
	}
}

func NewWriterToBuffer(buffer []byte) *MyWriter {
	return &MyWriter{
		data: buffer,
		pos:  0,
	}
}

// Читает данные из буфера в слайс p и возвращает количество прочитанных байт.
// EOF возвращается, если весь буфер был прочитан.
func (r *MyReader) Read(p []byte) (n int, err error) {
	if r.pos >= len(r.data) {
		return 0, io.EOF
	}

	n = copy(p, r.data[r.pos:])
	r.pos += n

	return n, nil
}

// Записывает данные из слайса p в буфер и возвращает количество записанных байт.
func (w *MyWriter) Write(p []byte) (n int, err error) {
	if w.pos >= len(w.data) {
		return 0, io.ErrShortWrite
	}

	// Определяем, сколько можно записать в буфер
	toCopy := len(p)
	freeLen := len(w.data) - w.pos
	if toCopy > freeLen {
		toCopy = freeLen
		err = io.ErrShortWrite
	}

	// Записываем, сколько влезает в буфер
	copy(w.data[w.pos:], p[:toCopy])
	w.pos += toCopy

	return toCopy, err
}

func Run533() {
	// Создаём буфер для записи и чтения данных длиной 20 байт
	buffer := make([]byte, 20)

	// Создаём MyWriter для записи в buffer
	writer := NewWriterToBuffer(buffer)

	// Записываем данные в буфер
	dataToWrite := []byte("Hello, World!")
	n, err := writer.Write(dataToWrite)
	if err != nil {
		fmt.Println("Ошибка записи:", err)
	} else {
		fmt.Printf("Записано %d байт: %s\n", n, string(buffer[:n]))
	}

	// Создаём MyReader для чтения из buffer
	reader := NewReaderFromBuffer(buffer)

	// Читаем данные из буфера чанками
	readBuffer := make([]byte, 5)
	for {
		n, err = reader.Read(readBuffer)
		if err != nil {
			if err.Error() == "EOF" {
				fmt.Printf("%s|EOF", string(readBuffer[:n])) // Выводим остаток, если достигнут EOF
				break
			}
			fmt.Println("Ошибка чтения:", err)
			break
		}
		fmt.Printf("|%s", (string(readBuffer[:n])))
	}
}
