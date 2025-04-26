package solution

type Palindrome struct {
	Str string
}

func (p Palindrome) IsPalindrome() bool {
	left := 0
	right := len(p.Str) - 1

	for left < right {
		if p.Str[left] != p.Str[right] {
			return false
		}
		left++
		right--
	}
	return true
}
