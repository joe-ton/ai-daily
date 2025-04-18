package solution

import "errors"

type TwoSum struct {
	Nums   []int
	Target int
}

// Find two indices from nums whose values are equal to target
func (t TwoSum) Find() (indices []int, err error) {
	if len(t.Nums) < 2 {
		return nil, errors.New("Invalid count of nums")
	}
	// TODO: handle empty values in fields for TwoSum

	intToIdx := make(map[int]int) // tracker of complements

	for idx, num := range t.Nums {
		complement := t.Target - num
		if compIdx, found := intToIdx[complement]; found {
			return []int{compIdx, idx}, nil
		}
		intToIdx[num] = idx // update to tracker of complements
	}
	return nil, errors.New("Invalid indices")
}
