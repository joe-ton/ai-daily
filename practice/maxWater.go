package practice

func GetMaxWater(nums []int) int {
	resultWater := 0
	left := 0
	right := len(nums) - 1

	for left < right {
		width := right - left
		area := width * min(nums[left], nums[right])
		resultWater = max(area, resultWater)

		if nums[left] < nums[right] {
			left++
		} else {
			right--
		}
	}
	return resultWater
}
