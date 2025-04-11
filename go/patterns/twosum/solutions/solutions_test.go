package solutions

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
			name:     "TestCorrectResponse",
			nums:     []int{1, 2, 3, 4},
			target:   7,
			wantResp: []int{2, 3},
			wantErr:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := TwoSumSolver{Nums: tt.nums, Target: tt.target}
			gotResp, gotErr := s.FindTwoSum()

			if !errors.Is(gotErr, tt.wantErr) {
				t.Errorf("\n%v \nGot:  %v\nWant: %v\n", tt.name, gotErr, tt.wantErr)
			}
			if reflect.DeepEqual(gotResp, tt.wantErr) {
				t.Errorf("\n%v \nGot:  %v\nWant: %v\n", tt.name, gotResp, tt.wantResp)
			}
		})
	}
}
