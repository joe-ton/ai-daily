package solution

import "errors"

type TwoSum struct {
	Nums   []int
	Target int
}

// Find two indices whose values together equal to the target
func (t TwoSum) Find() ([]int, error) {
	if len(t.Nums) < 2 {
		return nil, errors.New("Invalid count from Nums")
	}

	intToIdx := make(map[int]int) // complement lookup

	for idx, num := range t.Nums {
		complement := t.Target - num
		if compIdx, found := intToIdx[complement]; found {
			return []int{compIdx, idx}, nil
		}
		intToIdx[num] = idx
	}
	return nil, errors.New("Invalid count from Nums")
}
