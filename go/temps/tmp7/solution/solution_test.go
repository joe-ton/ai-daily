package solution

import "testing"

func TestPalindrome(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  bool
	}{
		{
			name:  "Test True 121",
			input: 121,
			want:  true,
		},
		{
			name:  "Test False 122",
			input: 122,
			want:  false,
		},
		{
			name:  "Test False Less Than Zero",
			input: -1,
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsPalindrome(tt.input)
			if got != tt.want {
				t.Errorf("Got %v, Want %v", got, tt.want)
			}
		})
	}
}
