package solution

import "errors"

type TwoSum struct {
	Nums   []int // ex. [1, 2, 3, 4]
	Target int   // ex. 7
}

// Find two indices whose values equal together given target
func (t TwoSum) Find() ([]int, error) {
	if len(t.Nums) < 2 {
		return nil, errors.New("Invalid count of Nums")
	}
	// TODO: checking empty variables

	seen := make(map[int]int) // int to idx

	for idx, num := range t.Nums {
		complement := t.Target - num
		if compIdx, found := seen[complement]; found {
			return []int{compIdx, idx}, nil
		}
		seen[num] = idx
	}
	return nil, errors.New("Invalid indices from Nums")
}
