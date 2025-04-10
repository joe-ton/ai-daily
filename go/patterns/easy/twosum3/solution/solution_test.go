// solution_test.go

package solution

import (
	"errors"
	"reflect"
	"testing"
)

func TestTwoSumFind(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		wantResp []int // indices
		wantErr  error
	}{
		{
			name:     "Test1",
			nums:     []int{1, 2, 3, 4},
			target:   7,
			wantResp: []int{2, 3},
			wantErr:  nil,
		},
		{
			name:     "TestGuardClauseTwoFewNums",
			nums:     []int{1},
			target:   7,
			wantResp: nil,
			wantErr:  errors.New("Invalid count of nums"),
		},
		{
			name:     "TestInvalidIndices",
			nums:     []int{1, 2, 3, 4},
			target:   8,
			wantResp: nil,
			wantErr:  errors.New("Invalid indices"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := TwoSumSolver{
				Nums:   tt.nums,
				Target: tt.target,
			}
			gotResp, gotErr := s.Solve()
			if !errors.Is(gotErr, tt.wantErr) {
				t.Errorf("\nGot:  %v\nWant: %v\n", gotErr, tt.wantErr)
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("\nGot:  %v\nWant: %v\n", gotResp, tt.wantResp)
			}

		})
	}
}
