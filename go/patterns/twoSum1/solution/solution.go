package solution

import "errors"

type TwoSum struct {
	Nums   []int
	Target int
}

// Find indices based off of the nums whose values equal target
func (t TwoSum) Find() (indices []int, err error) {
	if len(t.Nums) < 2 {
		return nil, errors.New("Invalid count of nums")
	}
	// TODO: empty inputs

	// target = var1 + var2
	seen := make(map[int]int)

	for idx, num := range t.Nums {
		complement := t.Target - num
		if compIdx, found := seen[complement]; found {
			return []int{compIdx, idx}, nil // early exist
		}
		seen[num] = idx // update to tracker
	}
	// fallback exist
	return nil, errors.New("Invalid indices")
}
