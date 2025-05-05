package solution

import "errors"

var (
	ErrInvalidCount   = errors.New("Invalid count of nums")
	ErrInvalidIndices = errors.New("Invalid indices of Find()")
)

type TwoSum struct {
	Nums   []int
	Target int
}

// Find two indices whose values equal together to target
func (t TwoSum) Find() (indices []int, err error) {
	if len(t.Nums) < 2 {
		return nil, ErrInvalidCount
	}

	seen := make(map[int]int) // lookup, int to idx

	for idx, num := range t.Nums {
		complement := t.Target - num
		if compIdx, found := seen[complement]; found {
			return []int{compIdx, idx}, nil
		}
		seen[num] = idx
	}
	return nil, ErrInvalidIndices
}
