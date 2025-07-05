package solution

import "errors"

type TwoSum struct {
	Nums   []int
	Target int
}

func (t TwoSum) Find() ([]int, error) {
	if len(t.Nums) < 2 {
		return nil, errors.New("Invalid Count")
	}

	seen := make(map[int]int)

	for idx, num := range t.Nums {
		complement := t.Target - num
		if compIdx, found := seen[complement]; found {
			return []int{compIdx, idx}, nil
		}
		seen[num] = idx
	}
	return nil, errors.New("No solution")
}
