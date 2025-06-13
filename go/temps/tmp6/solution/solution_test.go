package solution

import "testing"

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		{
			name: "Test True",
			s:    "cinema",
			t:    "iceman",
			want: true,
		},
		{
			name: "Test False Length",
			s:    "Person",
			t:    "Test",
			want: false,
		},
		{
			name: "Test False Anagram",
			s:    "aaa",
			t:    "abc",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Is_Anagram(tt.s, tt.t)
			if got != tt.want {
				t.Errorf("Got %v, Want %v", got, tt.want)
			}
		})
	}
}
