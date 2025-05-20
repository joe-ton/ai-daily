package solution

import "errors"

type TwoSumFinder interface {
	Find(nums []int, target int) ([]int, error)
}

type TwoSum struct {
	Nums   []int
	Target int
}

func (t TwoSum) FindBrute() ([]int, error) {
	if len(t.Nums) < 2 {
		return nil, errors.New("Invalid count of nums")
	}

	for idx1, num1 := range t.Nums {
		for idx2, num2 := range t.Nums[idx1+1:] {
			total := num1 + num2
			if total == t.Target {
				return []int{idx1, idx1 + idx2 + 1}, nil
			}
		}
	}
	return nil, errors.New("Invalid indices from nums")
}

func (t TwoSum) FindOptimize() ([]int, error) {
	if len(t.Nums) < 2 {
		return nil, errors.New("Invalid count of nums")
	}

	seen := make(map[int]int)

	for idx, num := range t.Nums {
		complement := t.Target - num
		if compIdx, found := seen[complement]; found {
			return []int{compIdx, idx}, nil
		}
		seen[num] = idx
	}
	return nil, errors.New("Invalid indices of nums")
}
