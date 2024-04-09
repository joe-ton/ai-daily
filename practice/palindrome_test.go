package practice

import "testing"

func Test_isPalindrome(t *testing.T) {
	tests := []struct {
		text   string // contains letters, mixed upper/lower, numbers
		isTrue bool   // result
	}{
		{"aba", true},
		{"abc", false},
		{"121", true},
	}
	for _, test := range tests {
		result := IsPalindrome(test.text)
		if result != test.isTrue {
			t.Errorf("Result Wrong: input %s, want %t, got %t",
				test.text, test.isTrue, result)
		}
	}
}
