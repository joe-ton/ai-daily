package removeduplicates

func removeDuplicates2(nums []int) int {
	// Given sorted nums, return count of unique integers
	nextInsertionIndex := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[nextInsertionIndex] = nums[i]
			nextInsertionIndex++
		}
	}
	return nextInsertionIndex
}
