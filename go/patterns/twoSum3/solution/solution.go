package solution

import "errors"

type TwoSum struct {
	Nums   []int
	Target int
}

// Find two indices from Nums whose values equal to sum target
func (t TwoSum) Find() (indices []int, err error) {
	if len(t.Nums) < 2 {
		return nil, errors.New("Invalid count of nums")
	}
	// TODO: empty values

	intToIdx := make(map[int]int) // tracker for complement

	for idx, num := range t.Nums {
		complement := t.Target - num
		if compIdx, found := intToIdx[complement]; found {
			return []int{compIdx, idx}, nil
		}
	}

}
