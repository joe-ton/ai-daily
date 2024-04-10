package practice

import (
	"reflect"
	"testing"
)

func Test_GetTwoSum(t *testing.T) {
	tests := []struct {
		nums    []int // unordered
		target  int   // two ints from nums added to sum
		indices []int // result
	}{
		{[]int{1, 2, 3, 1}, 5, []int{1, 2}},
	}
	for _, test := range tests {
		result := GetTwoSum(test.nums, test.target)
		if !reflect.DeepEqual(result, test.indices) {
			t.Errorf("Wrong Result: nums=%v target=%d, want %v, got %v",
				test.nums, test.target, test.indices, result)
		}
	}
}
