package solution

import "errors"

type TwoSum struct {
	Nums   []int
	Target int
}

// Brute force n^2
func (t TwoSum) FindBrute() ([]int, error) {
	for idx, num := range t.Nums {
		for idx2, num2 := range t.Nums[idx+1:] {
			total := num + num2
			if total == t.Target {
				return []int{idx, idx + idx2 + 1}, nil
			}
		}
	}
	return nil, errors.New("Invalid indices")
}

func (t TwoSum) FindMap() ([]int, error) {
	seen := make(map[int]int)
	for idx, num := range t.Nums {
		complement := t.Target - num
		if compIdx, found := seen[complement]; found {
			return []int{compIdx, idx}, nil
		}
		seen[num] = idx
	}
	return nil, errors.New("Invalid indices")
}
