package palindrome

import "unicode"

func isPalindrome(str string) bool {
	left := 0
	right := len(str) - 1

	for left < right {
		for left < right && !isAlpha(str[left]) {
			left++
		}
		for left < right && !isAlpha(str[right]) {
			right--
		}

		if unicode.ToLower(rune(str[left])) !=
			unicode.ToLower(rune(str[right])) {
			return false
		}
		left++
		right--
	}
	return true
}

func isAlpha(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') ||
		(ch >= 'A' && ch <= 'Z')
}
