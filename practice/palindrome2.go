package practice

import "unicode"

func isPalindrome2(str string) bool {
	left := 0
	right := len(str) - 1

	for left < right {
		for left < right && !isAlphaNum(str[left]) {
			left++
		}
		for left < right && !isAlphaNum(str[right]) {
			right--
		}
		if unicode.ToLower(rune(str[left])) !=
			unicode.ToLower(rune(str[right])) {
			return isPalindrome(str, left+1, right) ||
				isPalindrome(str, left, right-1)
		}
		left++
		right--
	}
	return true
}

func isPalindrome(str string, left, right int) bool {
	for left < right {
		if str[left] != str[right] {
			left++
			right--
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
