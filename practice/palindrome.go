package practice

import "unicode"

func IsPalindrome(str string) bool {
	left := 0
	right := len(str) - 1

	for left < right {
		// ignore non-important characters
		for left < right && !isAlphaNum(str[left]) {
			left++
		}
		for left < right && !isAlphaNum(str[right]) {
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

func isAlphaNum(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') ||
		(ch >= 'A' && ch <= 'Z') ||
		(ch >= '0' && ch <= '9')
}
