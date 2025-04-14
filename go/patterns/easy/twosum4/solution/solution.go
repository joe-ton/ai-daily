package solution

import "errors"

var (
	ErrNoValues  = errors.New("Invalid input")
	ErrCount     = errors.New("Invalid count from Nums")
	ErrNoIndices = errors.New("No Indices found")
)

type TwoSum struct {
	Nums   *[]int
	Target *int
}

// Find two indices from Nums whose values equal to target
func (i *TwoSum) Find() (indices []int, err error) {
	if i.Nums == nil || i.Target == nil {
		return nil, ErrNoValues
	}

	if len(*i.Nums) < 2 {
		return nil, ErrCount
	}

	intToIdx := make(map[int]int)

	for idx, num := range *i.Nums {
		complement := *i.Target - num
		if compIdx, found := intToIdx[complement]; found {
			return []int{compIdx, idx}, nil
		}
		intToIdx[num] = idx
	}
	return nil, ErrNoIndices
}
