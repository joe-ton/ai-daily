package practice

import "testing"

func Test_GetMaxWater(t *testing.T) {
	tests := []struct {
		nums     []int // verticals
		maxWater int   // result, area of max water
	}{
		{[]int{1, 2, 3}, 2},
	}
	for _, test := range tests {
		result := GetMaxWater(test.nums)
		if result != test.maxWater {
			t.Errorf("Wrong Result: nums=%v, want %d, got %d",
				test.nums, test.maxWater, result)
		}
	}
}
