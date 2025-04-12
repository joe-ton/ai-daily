package solution

import "errors"

type TwoSumInput struct {
	Nums   []int
	Target int
}

func (i TwoSumInput) FindTwoSum() ([]int, error) {
	if len(i.Nums) < 2 {
		return nil, errors.New("Invalid count of nums")
	}

	intToIdx := make(map[int]int)

	for idx, num := range i.Nums {
		complement := i.Target - num
		if compIdx, found := intToIdx[complement]; found {
			return []int{compIdx, idx}, nil
		}
		intToIdx[num] = idx
	}
	return nil, errors.New("Invalid indices from nums")
}
