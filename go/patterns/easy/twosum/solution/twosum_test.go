// twosum_test.go

package solution

import (
	"errors"
	"reflect"
	"testing"
)

// immutable: do not change
var (
	nums        []int = []int{1, 2, 3, 4}
	nums2       []int = []int{1}
	numsEmpty   []int = []int{}
	target      int   = 7
	targetEmpty *int  = nil
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name     string
		nums     *[]int
		target   *int
		wantResp []int
		wantErr  error
	}{
		{
			name:     "TestComplete",
			nums:     &nums,
			target:   &target,
			wantResp: []int{2, 3},
			wantErr:  nil,
		},
		{
			name:     "TestNoMatchFound",
			nums:     &[]int{1, 2, 5},
			target:   intPtr(10),
			wantResp: nil,
			wantErr:  ErrInvalidIndices,
		},
		{
			name:     "TestEmptyNums",
			nums:     &numsEmpty,
			target:   &target,
			wantResp: nil,
			wantErr:  ErrInvalidCount,
		},
		{
			name:     "TestNilTarget",
			nums:     &nums,
			target:   targetEmpty,
			wantResp: nil,
			wantErr:  ErrEmptyInput,
		},
		{
			name:     "TestCountIssue",
			nums:     &nums2,
			target:   &target,
			wantResp: nil,
			wantErr:  ErrInvalidCount,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := TwoSumInput{Nums: tt.nums, Target: tt.target}
			gotResp, gotErr := input.Find()
			if !errors.Is(gotErr, tt.wantErr) {
				t.Errorf("\nName: %v\nGot:  %v\nWant: %v",
					tt.name, gotErr, tt.wantErr)
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("Got %v, Want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func intPtr(i int) *int { return &i }
