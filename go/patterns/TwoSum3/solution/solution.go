package solution

import "errors"

type TwoSum struct {
	Nums   []int
	Target int
}

// Find two indices from Nums whose values together equal target
func (t TwoSum) Find() (indices []int, err error) {
	if len(t.Nums) < 2 {
		return nil, errors.New("Invalid count of Nums")
	}
}
