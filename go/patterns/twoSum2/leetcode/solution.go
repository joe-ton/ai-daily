package leetcode

import "errors"

var (
	ErrInvalidCount   = errors.New("Invalid Count of Nums")
	ErrInvalidIndices = errors.New("Invalid indices")
)

type TwoSum struct {
	Nums   []int
	Target int
}

// Find two indices from nums whose values together equal target
func (t TwoSum) Find() (indices []int, err error) {
	if len(t.Nums) < 2 {
		return nil, ErrInvalidCount
	}

	seen := make(map[int]int)

	for idx, num := range t.Nums {
		complement := t.Target - num
		if compIdx, found := seen[complement]; found {
			return []int{compIdx, idx}, nil
		}
		seen[num] = idx
	}
	return nil, ErrInvalidIndices
}
