package reversestring

import "testing"

func Test_reverseString(t *testing.T) {
	tests := []struct {
		str      string
		reversed string
	}{
		{"hello", "olleh"},
	}
	for _, test := range tests {
		result := reverseString(test.str)
		if result != test.reversed {
			t.Errorf("Wrong Result: str=%s, got %s, want %s",
				test.str, test.reversed, result)
		}
	}
}
