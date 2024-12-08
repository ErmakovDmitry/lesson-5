package homework

import (
	"errors"
	"fmt"
	"io"
)

type myReader struct {
	data []byte
	pos  int
}

// NewReaderFromBuffer создает новый myReader из переданного буфера.
func newReaderFromBuffer(buffer []byte) *myReader {
	return &myReader{data: buffer}
}

// Read читает данные из буфера в слайс p и возвращает количество прочитанных байт.
// EOF будет возвращён если весь буфер был прочитан.
func (reader *myReader) read(p []byte) (n int, err error) {

	// if len(p) < 1 {
	// 	// return 0, errors.New("размер буфера не может быть меньше 1")
	// 	return 0, fmt.Errorf("размер буфера не может быть меньше 1")
	// }

	if reader.pos >= len(reader.data) {
		return 0, io.EOF
	}

	n = copy(p, reader.data[reader.pos:])
	reader.pos += n

	// Проверка на конец данных
	if reader.pos >= len(reader.data) {
		err = io.EOF
	}

	return n, err
}

func Run532() {
	src_byte_slice := []byte("Hello, World!\nHere is another line.")
	fmt.Println("Исходная строка:", string(src_byte_slice[:]))
	fmt.Println()

	// Размер буфера
	bufSize := 0
	buf := make([]byte, bufSize)

	reader := newReaderFromBuffer(src_byte_slice)

	fmt.Println("Чтение из буфера:", string(buf[:]))
	for {
		n, err := reader.read(buf)
		if err != nil {
			if errors.Is(err, io.EOF) {
				fmt.Println("buf with EOF:", string(buf[:n]))
				break
			}
			fmt.Println("Error reading from buffer:", err)
			return
		}
		fmt.Println("buf without EOF:", string(buf[:n]))
	}
}
