package practice

func getTwoSum(nums []int, target int) []int {
	// counter -- array (slices), hashmaps// access
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
