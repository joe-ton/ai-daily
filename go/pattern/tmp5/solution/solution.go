package solution

import "errors"

type TwoSum struct {
	nums   []int
	target int
}

func NewTwoSum(nums []int, target int) *TwoSum {
	return &TwoSum{
		nums:   nums,
		target: target,
	}
}

// Find two indices whose values from the nums that equal given target
func (t *TwoSum) Find() ([]int, error) {
	if len(t.nums) < 2 {
		return nil, errors.New("Invalid count of nums")
	}

	seen := make(map[int]int) // int to Idx, tracker of compliment

	for idx, num := range t.nums {
		complement := t.target - num
		if compIdx, found := seen[complement]; found {
			return []int{compIdx, idx}, nil
		}
		seen[num] = idx
	}
	return nil, errors.New("Invalid indices from nums")
}
