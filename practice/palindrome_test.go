package practice

import "testing"

func Test_isPalindrome(t *testing.T) {
	tests := []struct {
		text   string
		isTrue bool // result
	}{
		{"abc", true},
		{"abcc", false},
		{"121", true},
	}
	for _, test := range tests {
		result := IsPalindrome(test.text)
		if result != test.isTrue {
			t.Errorf("Error result: insert %s, want %t, got %t",
				test.text, test.isTrue, result)
		}

	}
}
