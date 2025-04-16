package solution

type TwoSum struct {
	Nums   []int
	Target int
}

// Find two indices from nums whose values are equal to target
func (t TwoSum) Find() (indices []int, err error) {
	if len(t.Nums) < 2 {
		return nil, err
	}
    intToIdx := make(map[int]int)

    for idx, num := range t.Nums {
        if compIdx, num := 
    }
}
