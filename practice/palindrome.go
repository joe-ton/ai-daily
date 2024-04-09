package practice

import "unicode"

func IsPalindrome(text string) bool {
	left := 0
	right := len(text)

	for left < right {
		for left < right && !isAlpaNum(text[left]) {
			left++
		}
		for left < right && !isAlpaNum(text[right]) {
			right--
		}

		if unicode.ToLower(rune(text[left])) !=
			unicode.ToLower(rune(text[right])) {
			return false
		}
		left++
		right--
	}
	return true
}

func isAlpaNum(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') ||
		(ch >= 'A' && ch <= 'Z') ||
		(ch >= '0' && ch <= '9')
}
