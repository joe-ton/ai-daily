package solution

import "errors"

type TwoSum struct {
	Nums   []int
	Target int
}

func (t TwoSum) Find() (indices []int, err error) {
	if len(t.Nums) < 2 {
		return nil, errors.New("Invalid count of Nums")
	}

	seen := make(map[int]int) // int to idx, lookup for complements

	for idx, num := range t.Nums {
		complement := t.Target - num
		if compIdx, found := seen[complement]; found {
			return []int{compIdx, idx}, nil
		}
		seen[num] = idx
	}
	return nil, errors.New("Invalid indices of Nums")
}
