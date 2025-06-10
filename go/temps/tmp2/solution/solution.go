package solution

import "errors"

// Determine given numbers if values are duplicate
// If true, return true
func Is_duplicate(nums []int) (bool, error) {
	if len(nums) < 2 {
		return false, errors.New("Invalid count of nums")
	}
	seen := make(map[int]struct{}, len(nums)) // hashset; previous values

	for _, val := range nums {
		if _, dup := seen[val]; dup {
			return true, nil
		}
		seen[val] = struct{}{} // update the seen, previous values
	}

	return false, errors.New("Invalid result")
}
