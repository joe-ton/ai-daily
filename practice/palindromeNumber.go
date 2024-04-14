package practice

func IsPalindromeNumber(x int) bool {
	// Negative numbers are not palindromes
	if x < 0 {
		return false
	}

	// Reverse the number
	reversed := 0
	num := x
	for num != 0 {
		reversed = reversed*10 + num%10
		num /= 10
	}

	// Check if the reversed number is equal to the original number
	return x == reversed
}
