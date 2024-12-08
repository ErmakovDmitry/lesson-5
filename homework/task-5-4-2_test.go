package homework

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestJson(t *testing.T) {
	cases := map[string]struct {
		data    string
		example Example
		err     string
	}{
		"empty": {
			data: "{}",
		},
		"A": {
			data:    "{\"A\": 123}",
			example: Example{A: 123},
		},
		"AB": {
			data:    "{\"A\": 123, \"B\": \"test\"}",
			example: Example{A: 123, B: "test"},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			res, err := jsonToStruct([]byte(tc.data))
			if tc.err != "" {
				require.Error(t, err, "test")
			}
			require.NotNil(t, res)
			require.Equal(t, tc.example, *res)
		})
	}
}
