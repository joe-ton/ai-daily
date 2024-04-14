package practice

import "testing"

func Test_IsPalindromeNumber(t *testing.T) {
	testCases := []struct {
		num      int
		expected bool
	}{
		{121, true},
		{123, false},
	}
	for _, test := range testCases {
		result := IsPalindromeNumber(test.num)
		if result != test.expected {
			t.Errorf("Wrong Result: num=%d, want %t, got %t",
				test.num, test.expected, result)
		}
	}
}
