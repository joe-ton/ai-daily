package practice

import "testing"

func TestMyInterface(t *testing.T) {
	testCases := []struct {
		name     string
		input    MyInterface
		expected int
	}{
		{
			name:     "Real Implementation",
			input:    MyImplementation{},
			expected: 42,
		},
	}
	for _, test := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.input.Method1()
		}
	}
}
