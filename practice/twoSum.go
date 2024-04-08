package practice

func getTwoSum(nums []int, target int) []int {
	intToIndex := make(map[int]int) // previous ints

	for idx, num := range nums {
		complement := target - num
		if key, ok := intToIndex[complement]; ok {
			return []int{key, idx}
		}
		intToIndex[num] = idx
	}
	return []int{}
}
