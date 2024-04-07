package twoSum

import (
	"reflect"
	"testing"
)

func Test_getTwoSum(t *testing.T) {
	tests := []struct {
		nums        []int // unordered integers
		target      int   // sum of two integers
		twoIntegers []int // result
	}{
		{[]int{1, 2, 3, 1}, 5, []int{1, 2}},
		{[]int{1, 4, 5, 6}, 11, []int{2, 3}},
	}
	for _, test := range tests {
		result := getTwoSum(test.nums, test.target)
		if !reflect.DeepEqual(result, test.twoIntegers) {
			t.Errorf("Got %v, want %v", result, test.twoIntegers)
		}
	}
}
