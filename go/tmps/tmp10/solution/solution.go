package solution

import "errors"

// Determine if given values have duplicate integers
// Return if true
func Is_duplicate(nums []int) (bool, error) {
	// hash-set; tracker of previous values
	seen := make(map[int]struct{}, len(nums))

	// iterate, values
	for _, val := range nums {
		// find if current val in hashset
		if _, dup := seen[val]; dup {
			return true, nil
		}
		// remember value
		seen[val] = struct{}{}
	}
	return false, errors.New("Invalid result")
}
