package main

import "fmt"

<<<<<<< HEAD
func main() {
}

func getTwoSum(target int, integers []int) []int {
    numMap := make(map[int]int)
=======
func TwoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)

	for idx, num := range nums { // Corrected loop variable names
		complement := target - num
		if index, ok := hashMap[complement]; ok {
			return []int{index, idx} // Corrected the order of indices in the return statement
		}
		hashMap[num] = idx
	}
	return []int{}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	res := TwoSum(nums, target)
	fmt.Println(res) // Output: [0, 1]
>>>>>>> 3ad8269796a151c3a9ebc2505cac760ba15aca3d
}
