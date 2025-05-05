package solution

import (
	"errors"
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		wantResp []int // indices
		wantErr  error
	}{
		{
			name:     "Test Correct",
			nums:     []int{1, 2, 3, 4},
			target:   7,
			wantResp: []int{2, 3}, // indices
			wantErr:  nil,
		},
		{
			name:     "Test ErrInvalidCount",
			nums:     []int{1},
			target:   7,
			wantResp: nil, // indices
			wantErr:  ErrInvalidCount,
		},
		{
			name:     "Test ErrInvalidIndices",
			nums:     []int{1, 2, 3, 4},
			target:   100,
			wantResp: nil, // indices
			wantErr:  ErrInvalidIndices,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			twoSum := TwoSum{Target: tt.target, Nums: tt.nums}
			gotResp, gotErr := twoSum.Find()

			if !errors.Is(gotErr, tt.wantErr) {
				t.Errorf("Got %v, want %v", gotErr, tt.wantErr)
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("Got %v, want%v", gotResp, tt.wantResp)
			}
		})
	}
}
