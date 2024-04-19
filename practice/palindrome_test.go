package practice

import "testing"

func Test_isPalindrome(t *testing.T) {
	testCases := []struct {
		str  string
		want bool // expected
	}{
		{"abc", false},
		{"aba", true},
	}
	for _, test := range testCases {
		result := IsPalindrome(test.str)
		if result != test.want {
			t.Errorf("Wrong Result: str=%s, want %t, got %t",
				test.str, test.want, result)
		}
	}
}
