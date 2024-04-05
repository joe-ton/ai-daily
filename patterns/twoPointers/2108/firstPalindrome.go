package firstPalindrome

import "unicode"

func isFirstPalindrome(words []string) string {

	for _, word := range words {
		if isPalindrome(word) {
			return word
		}
	}
	// if no palindrome
	return ""
}

func isPalindrome(word string) bool {
	left := 0
	right := len(word) - 1

	for left < right {
		for left < right && isAlphaNum(word[left]) {
			left++
		}
		for left < right && isAlphaNum(word[right]) {
			right--
		}

		if unicode.ToLower(rune(word[left])) !=
			unicode.ToLower(rune(word[right])) {
			return false
		}
		left++
		right--
	}
	return true
}

func isAlphaNum(character byte) bool {
	ch := character

	return (ch >= 'a' && ch <= 'z') ||
		(ch >= 'A' && ch <= 'Z') ||
		(ch >= '0' && ch <= '9')

}
