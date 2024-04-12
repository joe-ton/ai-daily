package practice

func GetTwoSum(nums []int, target int) []int {
	// From integers get two numbers to add to target, return the indices

	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums); j++ {
			sum := nums[i] + nums[j]
			if sum == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

// func GetTwoSum(nums []int, target int) []int {
// 	// From integers get two numbers to add to target, return the indices
// 	intToIndex := make(map[int]int)

// 	for idx, num := range nums {
// 		complement := target - num
// 		if pos, ok := intToIndex[complement]; ok {
// 			return []int{pos, idx}
// 		}
// 		intToIndex[num] = idx
// 	}
// 	return []int{}
// }
