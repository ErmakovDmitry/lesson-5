package homework

import (
	"errors"
	"io"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRead532(t *testing.T) {
	cases := map[string]struct {
		data     string
		bufSize  int
		bufItems []string
	}{
		"empty": {
			data: "",
		},
		"one char": {
			data:    "A",
			bufSize: 1,
			bufItems: []string{
				"A",
			},
		},
		"many chars, equal sized bufs": {
			data:    "ABC1ABC2ABC3",
			bufSize: 4,
			bufItems: []string{
				"ABC1", "ABC2", "ABC3",
			},
		},
		"many chars, non equal sized bufs": {
			data:    "ABC1ABC2A3",
			bufSize: 4,
			bufItems: []string{
				"ABC1", "ABC2", "A3",
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			data := []byte(tc.data)
			reader := NewReaderFromBuffer(data)

			bufIndex := 0
			for {
				buf := make([]byte, tc.bufSize)
				n, err := reader.Read(buf)
				require.LessOrEqual(t, n, tc.bufSize, "Read too many bytes")
				if err != nil && err != io.EOF {
					require.NoError(t, err)
				}
				if errors.Is(err, io.EOF) {
					break
				}
				require.Equal(t, tc.bufItems[bufIndex], string(buf[:n]), "Part of read buf does not match")
				bufIndex++
			}
		})
	}
}
