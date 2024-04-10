package reverseString

func reverseString(str string) string {
	word := []byte(str)
	left := 0
	right := len(str) - 1

	for left < right {
		word[left], word[right] = word[right], word[left]
		left++
		right--
	}
	return string(word)
}
