package palindrome

import "testing"

func Test_isPalindrome(t *testing.T) {
	tests := []struct {
		str string
		is  bool
	}{
		{"racecar", true},
	}

	for _, test := range tests {
		result := isPalindrome(test.str)
		if result != test.is {
			t.Errorf("want %s got %t", test.str, result)
		}
	}
}
