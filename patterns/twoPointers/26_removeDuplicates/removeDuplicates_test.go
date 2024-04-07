package removeduplicates

import "testing"

func Test_removeDuplicates(t *testing.T) {
	tests := []struct {
		nums        []int // sorted already
		countUnique int   // result
	}{
		{[]int{1, 2, 2, 3}, 3},
	}
	for _, test := range tests {
		result := removeDuplicates(test.nums)
		if result != test.countUnique {
			t.Errorf("Got %d, want %d", result, test.countUnique)
		}
	}
}
