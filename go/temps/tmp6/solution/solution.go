package solution

// Determines of both given strings are anagrams.
// If true, return true
func Is_Anagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var freq [26]int

	for idx := range s {
		freq[s[idx]-'a']++
		freq[t[idx]-'a']--
	}

	for idx := range freq {
		char := freq[idx]
		if char != 0 {
			return false
		}
	}
	return true
}
