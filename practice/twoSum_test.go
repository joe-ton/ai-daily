package practice

import (
	"reflect"
	"testing"
)

func TestGetTwoSum(t *testing.T) {
	tests := []struct {
		nums    []int // unordered integers
		target  int   // two ints from num to equal target
		indices []int // result
	}{
		{[]int{1, 2, 3, 1}, 5, []int{1, 2}},
		{[]int{1, 20, 30, 40}, 70, []int{2, 3}},
	}

	for _, test := range tests {
		result := GetTwoSum(test.nums, test.target)
		if !reflect.DeepEqual(result, test.indices) {
			t.Errorf("Wrong Result: nums=%v target=%d, want %v, got %v",
				test.nums, test.target, test.indices, result)
		}
	}
}
