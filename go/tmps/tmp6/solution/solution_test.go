package solution

import (
	"errors"
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name    string
		nums    []int // list of integers
		target  int
		want    []int
		wantErr error
	}{
		{
			name:    "Test Complete",
			nums:    []int{1, 2, 3, 4}, // list of integers
			target:  7,
			want:    []int{2, 3},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			twoSum := TwoSum{Nums: tt.nums, Target: tt.target}
			got, err := twoSum.Find()

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Got %v, Want %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Got %v, Want %v", got, tt.want)
			}
		})
	}
}
