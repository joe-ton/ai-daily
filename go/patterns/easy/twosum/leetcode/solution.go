package leetcode

import "errors"

type TwoSum struct {
	Nums   []int
	Target int
}

// Find two indices from nums whose values equal target
func (t TwoSum) Find() (indices []int, err error) {
	if len(t.Nums) < 2 {
		return nil, errors.New("Invalid errors")
	}
	// TODO: TwoSum

	intToIdx := make(map[int]int) // store complements

	for idx, num := range t.Nums {
		complement := t.Target - num
		if compIdx, found := intToIdx[complement]; found {
			return []int{compIdx, idx}, nil
		}
		intToIdx[num] = idx
	}
	return nil, errors.New("Invalid indices")
}
