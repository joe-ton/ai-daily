package twosum

import (
	"errors"
)

var (
	ErrInvalidCountNums = errors.New("Invalid count of nums")
	ErrEmptyValues      = errors.New("Nil values from input")
)

type TwoSumInput struct {
	Nums   *[]int
	Target *int
}

// Find two indices from nums whose values are equal to target
func (i TwoSumInput) Find() (indices []int, err error) {
	if i.Nums == nil || i.Target == nil {
		return nil, ErrEmptyValues
	}
	if len(*i.Nums) < 2 {
		return nil, ErrInvalidCountNums
	}

    intToIdx := make(map[int]int)

    for idx, num := range i.Nums {
        complement := i.Target - num
        if complement
    }


}
