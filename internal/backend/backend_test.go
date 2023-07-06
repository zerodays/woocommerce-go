package backend

import (
	"bytes"
	"io"
	"testing"
)

func TestFilterReader(t *testing.T) {
	cases := []struct {
		name     string
		value    io.ReadCloser
		expected []byte
	}{
		{
			name:     "empty",
			value:    io.NopCloser(bytes.NewReader([]byte{})),
			expected: []byte{},
		},
		{
			name:     "trivial",
			value:    io.NopCloser(bytes.NewReader([]byte{1})),
			expected: []byte{1},
		},
		{
			name:     "trivial empty",
			value:    io.NopCloser(bytes.NewReader([]byte{0})),
			expected: []byte{},
		},
		{
			name:     "complex",
			value:    io.NopCloser(bytes.NewReader([]byte{0, 0, 0, 0, 1, 2, 3, 0, 4, 5})),
			expected: []byte{1, 2, 3, 4, 5},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			f := &filterReader{c.value}
			data, err := io.ReadAll(f)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if !bytes.Equal(data, c.expected) {
				t.Fatalf("expected %v, got %v", c.expected, data)
			}
		})
	}
}
