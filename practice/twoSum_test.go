package practice

import (
	"reflect"
	"testing"
)

func Test_getTwoSum(t *testing.T) {
	testCases := []struct {
		nums     []int
		target   int
		expected []int
	}{
		{[]int{1, 2, 3, 1}, 5, []int{1, 2}},
	}
	for _, test := range testCases {
		result := getTwoSum(test.nums, test.target)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Wrong Result: nums=%v target=%d, want %v, got %v",
				test.nums, test.target, test.expected, result)
		}
	}
}
