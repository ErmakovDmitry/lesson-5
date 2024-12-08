package homework

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReverseRead(t *testing.T) {
	cases := map[string]struct {
		data     string
		expected []string
	}{
		"empty": {
			data:     "",
			expected: nil,
		},
		"couple lines": {
			data:     "first line\nsecond line",
			expected: []string{"second line", "first line"},
		},
		"many lines": {
			data:     "first line\nsecond line\nthird line",
			expected: []string{"third line", "second line", "first line"},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			res, err := reverseReader(strings.NewReader(tc.data))
			require.NoError(t, err)
			require.Equal(t, tc.expected, res)
		})
	}
}
