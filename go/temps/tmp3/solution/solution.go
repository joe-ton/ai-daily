package solution

import "errors"

func Is_duplicate(nums []int) (bool, error) {
	seen := make(map[int]struct{}, len(nums))

	for _, val := range nums {
		if _, dup := seen[val]; dup {
			return true, nil
		}
		seen[val] = struct{}{}
	}
	return false, errors.New("Invalid result")
}
