package containerWater

func GetMaxWater(heights []int) int {
	left := 0
	right := len(heights) - 1
	maxResult := 0

	for left < right {
		width := right - left
		area := width * min(heights[left], heights[right])
		maxResult = max(maxResult, area)

		if heights[left] < heights[right] {
			left++
		} else {
			right--
		}
	}
	return maxResult
}
