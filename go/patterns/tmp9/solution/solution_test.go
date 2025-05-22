package solution

import (
	"errors"
	"reflect"
	"testing"
)

func TestTwoSumBrute(t *testing.T) {
	tests := []struct {
		name    string
		nums    []int
		target  int
		want    []int // indices
		wantErr error
	}{
		{
			name:    "Test Complete",
			nums:    []int{1, 2, 3, 4},
			target:  7,
			want:    []int{2, 3}, // indices
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			twoSum := TwoSum{Nums: tt.nums, Target: tt.target}
			got, err := twoSum.FindBrute()
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Got %v, Want %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Got %v, Want %v", got, tt.want)
			}
		})
	}

}

func TestTwoSumMap(t *testing.T) {
	tests := []struct {
		name    string
		nums    []int
		target  int
		want    []int // indices
		wantErr error
	}{
		{
			name:    "Test Complete",
			nums:    []int{1, 2, 3, 4},
			target:  7,
			want:    []int{2, 3}, // indices
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			twoSum := TwoSum{Nums: tt.nums, Target: tt.target}
			got, err := twoSum.FindMap()
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Got %v, Want %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Got %v, Want %v", got, tt.want)
			}
		})
	}

}
