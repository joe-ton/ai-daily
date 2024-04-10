package practice

func GetTwoSum(nums []int, target int) []int {
	intToIndex := make(map[int]int) // previous

	for idx, num := range nums {
		complement := target - num
		if key, ok := intToIndex[complement]; ok {
			return []int{key, idx}
		}
		intToIndex[num] = idx
	}
	return []int{}
}
