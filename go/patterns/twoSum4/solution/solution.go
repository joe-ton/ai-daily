package solution

import "errors"

type TwoSum struct {
	Nums   []int
	Target int
}

// Find two indices from nums whose values together to target
func (t TwoSum) Find() (indices []int, err error) {
	if len(t.Nums) < 2 {
		return nil, errors.New("Invalid count of nums")
	}

	seen := make(map[int]int) // num to idx

	for idx, num := range t.Nums {
		complement := t.Target - num
		if compIdx, found := seen[complement]; found {
			return []int{compIdx, idx}, nil
		}
		seen[num] = idx
	}
	return nil, errors.New("Invalid indices")
}
