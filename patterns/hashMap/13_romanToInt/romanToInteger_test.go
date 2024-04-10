package romantointeger

import "testing"

func Test_romanToInteger(t *testing.T) {
	tests := []struct {
		str   string // roman numeral
		value int    // result, convert roman to integer
	}{
		{"III", 3},
		{"VI", 6},
	}
	for _, test := range tests {
		result := romanToInteger(test.str)
		if result != test.value {
			t.Errorf("Wrong Result: str=%s, want %d, got %d",
				test.str, test.value, result)
		}
	}
}
