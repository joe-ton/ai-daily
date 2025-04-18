package solution

import "errors"

type TwoSum struct {
	Nums   []int
	Target int
}

// Find two indices from nums to equal target
func (t TwoSum) Find() (indices []int, err error) {
	// count Nums
	if len(t.Nums) < 2 {
		return nil, errors.New("Invalid counts in nums")
	}
	// TODO: values?

	intToIdx := make(map[int]int) // tracker

	for idx, num := range t.Nums {
		complement := t.Target - num
		if compIdx, found := intToIdx[complement]; found {
			return []int{compIdx, idx}, nil
		}
		intToIdx[num] = idx
	}
	return nil, errors.New("Invalid indices")
}
