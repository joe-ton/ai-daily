package solution

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	seen := make(map[int]struct{}, len(nums))

	for _, val := range nums {
		seen[val] = struct{}{}
	}

	longest := 0

	for idx := range seen {
		if _, hasPrev := seen[idx-1]; hasPrev {
			continue
		}

		length := 1

		for cur := idx + 1; ; cur++ {
			if _, ok := seen[cur]; !ok {
				break
			}
		}
		if length < longest {
			longest = length
		}
	}
	return longest
}
