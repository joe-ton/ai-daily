package practice

import "testing"

func Test_reverseString(t *testing.T) {
	tests := []struct {
		str        string
		reverseStr string // result
	}{
		{"hello", "olleh"},
		{"abc", "cba"},
		{"greetings,", ",sgniteerg"},
	}
	for _, test := range tests {
		result := ReverseString(test.str)
		if result != test.reverseStr {
			t.Errorf("Wrong Result: str=%s, want %s, got %s",
				test.str, test.reverseStr, result)
		}
	}
}
