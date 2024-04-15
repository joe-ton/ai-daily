package practice

import "testing"

func Test_isPalindrome2(t *testing.T) {
	testCases := []struct {
		str      string
		expected bool
	}{
		{"abbca", true},
		{"abbcda", false},
	}
	for _, test := range testCases {
		result := isPalindrome2(test.str)
		if result != test.expected {
			t.Errorf("Wrong Result: str=%s, want %t, got %t",
				test.str, test.expected, result)
		}
	}
}
