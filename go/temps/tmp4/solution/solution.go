package solution

import "errors"

// Determine given numbers have duplicate values
// If true, return true
func Is_Duplicate(nums []int) (bool, error) {
	if len(nums) < 2 {
		return false, errors.New("Invalid count of Nums")
	}

	seen := make(map[int]struct{}, len(nums))

	for _, val := range nums {
		if _, dup := seen[val]; dup {
			return true, nil
		}
		seen[val] = struct{}{}
	}
	return false, errors.New("Invalid response")
}
