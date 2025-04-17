package solution

import "errors"

var (
	ErrInvalidNumsCount = errors.New("Invalid given count of nums")
	ErrInvalidIndices   = errors.New("Invalid indices")
)

type TwoSum struct {
	Nums   []int
	Target int
}

// Find two indices from nums integers to equal to target
func (t TwoSum) Find() (indices []int, err error) {
	if len(t.Nums) < 2 {
		return nil, ErrInvalidNumsCount
	}
	// TODO: empty values, guard

	intToIdx := make(map[int]int) // tracker

	for idx, num := range t.Nums {
		complement := t.Target - num
		if compIdx, found := intToIdx[complement]; found {
			return []int{compIdx, idx}, nil
		}
		intToIdx[num] = idx
	}
	return nil, ErrInvalidIndices
}
