// twosum.go

package solution

import "errors"

var (
	ErrEmptyInput     = errors.New("Invalid Input")
	ErrInvalidCount   = errors.New("Not enough count for nums")
	ErrInvalidIndices = errors.New("Incorrect indices")
)

type TwoSumInput struct {
	Nums   *[]int
	Target *int
}

// Find two indices from nums whose values equal to target
func (i *TwoSumInput) Find() (indices []int, err error) {
	if i.Nums == nil || i.Target == nil {
		return nil, ErrEmptyInput
	}

	if len(*i.Nums) < 2 {
		return nil, ErrInvalidCount
	}

	intToIdx := make(map[int]int)

	for idx, num := range *i.Nums {
		complement := *i.Target - num
		if compIdx, found := intToIdx[complement]; found {
			return []int{compIdx, idx}, nil
		}
		intToIdx[num] = idx
	}
	return nil, ErrInvalidIndices
}
