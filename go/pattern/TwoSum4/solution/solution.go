package solution

import "errors"

type TwoSum struct {
	Nums   []int
	Target int
}

// Find two indices whose values together equal given target
func (t TwoSum) Find() ([]int, error) {
	if len(t.Nums) < 2 {
		return nil, errors.New("Invalid count of Nums")
	}
	// TODO: blanks, target

	seen := make(map[int]int) // lookup

	for idx, num := range t.Nums {
		complement := t.Target - num
		if compIdx, found := seen[complement]; found {
			return []int{compIdx, idx}, nil
		}
		seen[num] = idx
	}
	return nil, errors.New("Invalid indices from Nums")
}
