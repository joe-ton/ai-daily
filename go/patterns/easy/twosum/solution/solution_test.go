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
		wantResp []int
		wantErr  error
	}{
		{
			name:     "TestCorrect",
			nums:     []int{1, 2, 3, 4},
			target:   7,
			wantResp: []int{2, 3},
			wantErr:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := TwoSumInput{Nums: tt.nums, Target: tt.target}
			gotResp, gotErr := input.FindTwoSum()
			if !errors.Is(gotErr, tt.wantErr) {
				t.Errorf("Got %v, Want %v", gotErr, tt.wantErr)
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("Got %v, Want %v", gotResp, tt.wantResp)
			}
		})
	}
}
