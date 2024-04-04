package containerWater

import "testing"

func Test_getMaxWater(t *testing.T) {
	tests := []struct {
		verticals []int
		maxArea   int
	}{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
	}

	for _, test := range tests {
		result := GetMaxWater(test.verticals)
		if result != test.maxArea {
			t.Errorf("got %v, want %d", result, test.maxArea)
		}
	}
}
