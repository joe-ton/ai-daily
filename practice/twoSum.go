package practice

func GetTwoSum(nums []int, target int) []int {
	// Given unordered integers, find two integers to add to target
	// Return the indices of the two integers

	intToIndex := make(map[int]int)

	for idx, num := range nums {
		complement := target - num
		if key, ok := intToIndex[complement]; ok {
			return []int{key, idx}
		}
		intToIndex[num] = idx
	}
	return []int{}
}
