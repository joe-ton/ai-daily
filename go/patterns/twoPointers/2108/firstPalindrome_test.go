package firstPalindrome

import "testing"

func Test_isFirstPalindrome(t *testing.T) {
	tests := []struct {
		sliceOfWords    []string
		firstPalindrome string
	}{
		{[]string{"person", "racecar"}, "racecar"},
	}

	for _, test := range tests {
		result := isFirstPalindrome(test.sliceOfWords)
		if result != test.firstPalindrome {
			t.Errorf("Want %s, got %s", test.firstPalindrome, result)
		}
	}
}

func Test_isPalindrome(t *testing.T) {
	tests := []struct {
		word     string
		response bool
	}{
		{"word", false},
		{"racecar", true},
		{"121", true},
	}
	for _, test := range tests {
		result := isPalindrome(test.word)
		if result != test.response {
			t.Errorf("Word='%s', got %t, want %t", test.word,
				result, test.response)
		}
	}
}
