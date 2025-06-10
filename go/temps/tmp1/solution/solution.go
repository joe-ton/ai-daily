package solution

import "errors"

// Determine if given numbers have duplicate values
// If true, return true. Otherwise false
func Is_duplicate(nums []int) (bool, error) {
	// hash-set
	seen := make(map[int]struct{}, len(nums))

	// iterate through
	for _, val := range nums {
		if _, dup := seen[val]; dup {
			return true, nil
		}
		seen[val] = struct{}{}
	}
	return false, errors.New("Invalid bool")
}
