package solution

import "errors"

var (
	ErrInvalidNumsCount = errors.New("Invalid count of nums")
)

type TwoSum struct {
	Nums   []int
	Target int
}

func (t TwoSum) Find() (indices []int, err error) {
	if len(t.Nums) < 2 {
		return nil, ErrInvalidNumsCount
	}
	// TODO: errs - empty values

	intToIdx := make(map[int]int)

	for idx, num := range t.Nums {
		complement := t.Target - num
		if compIdx, found := intToIdx[complement]; found {
			return []int{compIdx, idx}, nil
		}
		intToIdx[num] = idx
	}
	return nil, errors.New("Invalid indices")
}
