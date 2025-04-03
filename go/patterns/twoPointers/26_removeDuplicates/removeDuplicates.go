package removeduplicates

func removeDuplicates(nums []int) int {
	prev := nums[0]
	left := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] != prev {
			nums[left] = nums[i]
			left++
		}
		prev = nums[i]
	}
	return left
}
