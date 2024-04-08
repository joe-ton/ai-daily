package practice

import (
	"reflect"
	"testing"
)

func Test_getTwoSum(t *testing.T) {
	tests := []struct {
		nums    []int // unordered
		target  int   // sum of two ints
		indices []int // result, indices of the two ints
	}{
		{[]int{1, 2, 3, 1}, 5, []int{1, 2}},
		{[]int{1, 2, 3, 10}, 12, []int{1, 3}},
	}
	for _, test := range tests {
		result := getTwoSum(test.nums, test.target)
		if !reflect.DeepEqual(result, test.indices) {
			t.Errorf("Got %v, want %v", result, test.indices)
		}
	}
}
