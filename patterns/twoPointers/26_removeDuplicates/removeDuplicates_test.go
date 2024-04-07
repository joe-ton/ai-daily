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
		result1 := removeDuplicates(test.nums)
		result2 := removeDuplicates2(test.nums)
		if result1 != test.countUnique {
			t.Errorf("Got %d, want %d", result1, test.countUnique)
		}
		if result2 != test.countUnique {
			t.Errorf("Got %d, want %d", result2, test.countUnique)
		}
	}
}
