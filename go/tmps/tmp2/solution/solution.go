package solution

import "errors"

type TwoSum struct {
	Nums   []int
	Target int
}

func (t TwoSum) Find() ([]int, error) {
	if len(t.Nums) < 2 {
		return nil, errors.New("Invalid count")
	}

	for idx, num := range t.Nums {

	}

	return nil, errors.New("Invalid solution")
}
