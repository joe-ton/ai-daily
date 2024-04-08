package practice

import "testing"

func Test_getMaxWater(t *testing.T) {
	tests := []struct {
		nums     []int // verticals
		maxWater int   // result, maxArea of water
	}{
		{[]int{1, 2, 3}, 2},
	}
	for _, test := range tests {
		result := getMaxWater(test.nums)
		if result != test.maxWater {
			t.Errorf("Result incorrect: got %d, want %d",
				result, test.maxWater)
		}
	}
}
