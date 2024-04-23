package practice

import (
	"reflect"
	"testing"
)

func Test_getTwoSum(t *testing.T) {
	testCases := []struct {
		input_nums    []int
		input_target  int
		output_slices []int
	}{
		{[]int{1, 2, 3, 1}, 5, []int{1, 2}},
		{[]int{}, 5, []int{}},
	}
	for _, tc := range testCases {
		result := getTwoSum(tc.input_nums, tc.input_target)
		if !reflect.DeepEqual(tc.output_slices, result) {
			t.Errorf("Wrong Result: nums=%v target=%d, want %v, got %v",
				tc.input_nums, tc.input_target, tc.output_slices, result)
		}
	}
}
