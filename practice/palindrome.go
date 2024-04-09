package practice

import "unicode"

func IsPalindrome(text string) bool {
	left := 0
	right := len(text) - 1

	for left < right {
		// ignore non-alphanum characters
		for left < right && !isAlphaNum(text[left]) {
			left++
		}
		for left < right && !isAlphaNum(text[right]) {
			right--
		}
		// acutal comparison
		if unicode.ToLower(rune(text[left])) !=
			unicode.ToLower(rune(text[right])) {
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
