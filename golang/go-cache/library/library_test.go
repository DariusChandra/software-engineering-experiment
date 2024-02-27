package library

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{0, "hello world"},
		{30, "halo"},
		{42, "halo"}, // Add more test cases as needed
	}

	for _, test := range tests {
		result := HelloWorld(test.input)
		if result != test.expected {
			t.Errorf("HelloWorld(%d) returned %s, expected %s", test.input, result, test.expected)
		}
	}
}
