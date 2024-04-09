package practice

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		str    string // random text, need ToLower
		isTrue bool   // result
	}{
		{"aba", true},
		{"abc", false},
		{"121", true},
	}
	for _, test := range tests {
		result := IsPalindrome(test.str)
		if result != test.isTrue {
			t.Errorf("Wrong Result: str=%s, got %t, want %t",
				test.str, test.isTrue, result)
		}
	}
}
