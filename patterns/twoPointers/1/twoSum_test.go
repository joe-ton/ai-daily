package twoSum

import (
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	tests := []struct {
		nums    []int
		target  int
		indices []int
	}{
		{[]int{1, 2, 3, 1}, 5, []int{1, 2}},
	}

	for _, test := range tests {
		result := twoSum(test.nums, test.target)
		if !reflect.DeepEqual(result, test.indices) {
			t.Errorf("want %v, got %v", test.indices, result)
		}
	}

}
