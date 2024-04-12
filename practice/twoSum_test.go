package practice

import (
	"reflect"
	"testing"
)

func Test_GetTwoSum(t *testing.T) {
	tests := []struct {
		nums    []int
		target  int
		indices []int
	}{
		{[]int{1, 2, 3, 4}, 7, []int{2, 3}},
		{[]int{10, 2, 3, 4}, 12, []int{0, 1}},
	}
	for _, test := range tests {
		result := GetTwoSum(test.nums, test.target)
		if !reflect.DeepEqual(result, test.indices) {
			t.Errorf("Wrong Result: nums=%v target=%d, want %v, got %v",
				test.nums, test.target, test.indices, result)
		}
	}
}
