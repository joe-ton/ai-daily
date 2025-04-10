// solution.go

package solution

import "errors"

type TwoSumSolver struct {
	Nums   []int
	Target int
}

// Find two indices from nums to sum up to target
func (s TwoSumSolver) Solve() ([]int, error) {
	if len(s.Nums) < 2 {
		return nil, errors.New("Invalid count of nums")
	}

	seen := make(map[int]int) // int to index

	for idx, num := range s.Nums {
		complement := s.Target - num
		if compIdx, found := seen[complement]; found {
			return []int{compIdx, idx}, nil
		}
		seen[num] = idx
	}
	return nil, errors.New("Invalid indices")
}
