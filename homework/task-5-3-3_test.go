package homework

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRead533(t *testing.T) {
	cases := map[string]struct {
		data string
	}{
		"empty": {
			data: "",
		},
		"one char": {
			data: "A",
		},
		"many chars": {
			data: "ABCABCABC",
		},
		"long string": {
			data: strings.Repeat("ABC", 100),
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			data := []byte(tc.data)
			buffer := make([]byte, len(data))
			reader := NewReaderFromBuffer(data)
			writer := NewWriterToBuffer(buffer)

			n, err := io.Copy(writer, reader)
			require.NoError(t, err, "unexpected error")
			require.Equal(t, int64(len(data)), n)
			require.Equal(t, tc.data, string(writer.data))
		})
	}
}
