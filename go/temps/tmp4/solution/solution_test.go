package solution

import (
	"errors"
	"reflect"
	"testing"
)

func TestIsDuplicate(t *testing.T) {
	tests := []struct {
		name    string
		nums    []int
		want    bool
		wantErr error
	}{
		{
			name:    "Test complete",
			nums:    []int{1, 2, 3, 3},
			want:    true,
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Is_Duplicate(tt.nums)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Got %v, Want %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Got %v, Want %v", got, tt.want)
			}
		})
	}
}
