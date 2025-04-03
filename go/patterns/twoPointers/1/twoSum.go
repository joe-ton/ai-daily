package twoSum

func twoSum(nums []int, target int) []int {
	intToIndex := make(map[int]int)

	for idx, num := range nums {
		complement := target - num
		if pos, ok := intToIndex[complement]; ok {
			return []int{pos, idx}
		}
		intToIndex[num] = idx
	}
	return []int{}
}
