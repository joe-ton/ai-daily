package practice

import "unicode"

func IsPalindrome(str string) bool {
	left := 0
	right := len(str) - 1

	for left < right {
		for left < right && !isAlpaNum(str[left]) {
			left++
		}
		for left < right && !isAlpaNum(str[right]) {
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

func isAlpaNum(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') ||
		(ch >= 'A' && ch <= 'Z') ||
		(ch >= '0' && ch <= '9')
}
