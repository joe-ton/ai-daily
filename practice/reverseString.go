package reversestring

func reverseString(str string) string {
	byteStr := []byte(str)
	left := 0
	right := len(str) - 1

	for left < right {
		byteStr[left], byteStr[right] = byteStr[right], byteStr[left]
		left++
		right--
	}
	return string(byteStr)
}
