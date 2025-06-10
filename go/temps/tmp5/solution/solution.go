package solution

func is_anagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var freq [26]int

	for i := range s {
		freq[s[i]-'a']++
		freq[t[i]-'a']--
	}

	// verifcation step
	// scan each 26 counters
	// this to ensure that we have the counters as 0
	for _, val := range freq {
		if val != 0 {
			return false
		}
	}

	return true
}
