package reversewords

func reverseWords(s string) string {
	bytes := []byte(s)
	lastSpaceIndex := -1

	for i := 0; i <= len(bytes); i++ {
		if i == len(bytes) || bytes[i] == ' ' {
			startIndex := lastSpaceIndex + 1
			endIndex := i - 1
			reverse(bytes, startIndex, endIndex)
			lastSpaceIndex = i
		}
	}
	return string(bytes)
}

func reverse(bytes []byte, startIndex, endIndex int) {
	for startIndex < endIndex {
		bytes[startIndex], bytes[endIndex] = bytes[endIndex], bytes[startIndex]
		startIndex++
		endIndex--
	}
}
