package solution

import "errors"

type TwoSum struct {
	Nums   []int
	Target int
}

func (s *TwoSum) Find() ([]int, error) {
	if len(s.Nums) < 2 {
		return nil, errors.New("Invalid count of nums")
	}

	seen := make(map[int]int) // int to idx for complement

	for idx, num := range s.Nums {
		complement := s.Target - num
		if compIdx, found := seen[complement]; found {
			return []int{compIdx, idx}, nil
		}
		seen[num] = idx
	}
	return nil, errors.New("Invalid indices from nums")
}
