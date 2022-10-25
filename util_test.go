package woocommerce

import (
	"bytes"
	"strconv"
	"testing"
)

func TestInt_MarshalJSON(t *testing.T) {
	cases := []struct {
		value    Int
		expected []byte
	}{
		{
			value:    0,
			expected: []byte(`"0"`),
		},
		{
			value:    42,
			expected: []byte(`"42"`),
		},
		{
			value:    -69,
			expected: []byte(`"-69"`),
		},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual, err := c.value.MarshalJSON()
			if err != nil {
				t.Fatal(err)
			}

			if bytes.Compare(actual, c.expected) != 0 {
				t.Fatalf("expected %s, got %s", string(c.expected), string(actual))
			}
		})
	}
}

func TestInt_UnmarshalJSON(t *testing.T) {
	cases := []struct {
		value       []byte
		expected    Int
		expectedErr bool
	}{
		{
			value:    []byte(`"0"`),
			expected: 0,
		},
		{
			value:    []byte(`"42"`),
			expected: 42,
		},
		{
			value:    []byte(`"-69"`),
			expected: -69,
		},
		{
			value:       []byte(`"not a number"`),
			expectedErr: true,
		},
		{
			value:       []byte(``),
			expectedErr: true,
		},
	}
	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			var actual Int
			err := actual.UnmarshalJSON(c.value)
			if c.expectedErr && err == nil {
				t.Fatal("expected error, got nil")
			} else if !c.expectedErr && err != nil {
				t.Fatal(err)
			}

			if actual != c.expected {
				t.Fatalf("expected %d, got %d", c.expected, actual)
			}
		})
	}
}
