package containerWater

func GetMaxWater(verticals []int) int {
	left := 0
	right := len(verticals) - 1
	maxResult := 0

	for left < right {
		width := verticals[left] - verticals[right]
		area := width * min(verticals[left], verticals[right])
		maxResult = max(maxResult, area)

		if verticals[left] < verticals[right] {
			left++
		} else {
			right--
		}
	}
	return maxResult
}
