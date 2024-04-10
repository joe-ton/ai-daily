package reversestring

import "testing"

func Test_reverseString(t *testing.T) {
	tests := []struct {
		text        string
		reverseText string // result
	}{
		{"hello", "olleh"},
	}
	for _, test := range tests {
		result := reverseString(test.text)
		if result != test.reverseText {
			t.Errorf("Wrong Result: text=%s, want %s, got %s",
				test.text, test.reverseText, result)
		}
	}
}
