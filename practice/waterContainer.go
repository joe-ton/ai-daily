package practice

func getMaxWater(nums []int) int {
	maxArea := 0
	left := 0
	right := len(nums) - 1

	for left < right {
		width := right - left
		area := width * min(nums[left], nums[right])
		maxArea = max(maxArea, area)

		if nums[left] < nums[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}
