package twoSum

func getTwoSum(nums []int, target int) []int {
	intToIndex := make(map[int]int) // previous integers

	for idx, num := range nums {
		complement := target - num
		if complementIdx, ok := intToIndex[complement]; ok {
			return []int{complementIdx, idx}
		}
		intToIndex[num] = idx
	}
	return []int{}
}
