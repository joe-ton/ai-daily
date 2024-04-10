package practice

func ReverseString(str string) string {
	strByte := []byte(str)
	left := 0
	right := len(str) - 1

	for left < right {
		strByte[left], strByte[right] = strByte[right], strByte[left]
		left++
		right--
	}
	return string(strByte)
}
