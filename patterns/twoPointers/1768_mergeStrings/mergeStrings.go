package mergestrings

func mergeStrings(word1 string, word2 string) string {
	result := ""
	i := 0
	for i = range word1 {
		if i == len(word2) {
			return result + word1[i:]
		}
		result += string(word1[i]) + string(word2[i])
	}
	return result + word2[i+1:]
}
