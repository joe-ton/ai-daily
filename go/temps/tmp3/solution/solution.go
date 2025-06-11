package solution

func is_anagram(s, t []int) bool {
	if len(s) != len(t) {
		return false
	}

	var freq [26]int

	for _, char := range s {
		freq[char-'a']++
		freq[char-'a']--
	}

	for _, char := range freq {
		if char == 0 {
			return true
		}
	}
	return false
}
