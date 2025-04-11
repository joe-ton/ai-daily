package solutions

import "errors"

type TwoSumSolver struct {
	Nums   []int
	Target int
}

// Find two indices whose values are sum to target
func (s TwoSumSolver) FindTwoSum() ([]int, error) {
	if len(s.Nums) < 2 {
		return nil, errors.New("Invalid count to nums")
	}
	seen := make(map[int]int) // num to idx

	for idx, num := range s.Nums {
		complement := s.Target - num
		if compIdx, found := seen[complement]; found {
			return []int{compIdx, idx}, nil
		}
		seen[num] = idx
	}
	return nil, errors.New("Invalid indices")
}
