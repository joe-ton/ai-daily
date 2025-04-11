package solution

import "errors"

type TwoSumInput struct {
	Nums   []int
	Target int
}

// Find two indices from two values from the given nums to add to target
func (i TwoSumInput) FindTwoSum() ([]int, error) {
	if len(i.Nums) < 2 {
		return nil, errors.New("Invalid count on nums")
	}

	seen := make(map[int]int) // int to idx

	for idx, num := range i.Nums {
		complement := i.Target - num

		if compIdx, found := seen[complement]; found {
			return []int{compIdx, idx}, nil
		}
		seen[num] = idx
	}
	return nil, errors.New("Invalid indices")
}
