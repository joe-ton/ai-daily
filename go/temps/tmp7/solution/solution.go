package solution

import "strconv"

// Given x int, determine if x is a palindrome.
// If palindrome, return true.  Otherwise return false.
func IsPalindrome(input int) bool {
	if input < 0 {
		return false
	}

	val := strconv.Itoa(input)

	leftIdx := 0
	rightIdx := len(val) - 1

	for leftIdx < rightIdx {
		leftVal := val[leftIdx]
		rightVal := val[rightIdx]

		if leftVal != rightVal {
			return false
		}

		leftIdx++
		rightIdx--
	}
	return true
}
