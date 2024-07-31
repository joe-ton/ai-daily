package main

import "fmt"

func getTwoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)

	for num, idx := range nums {
		complement := target - num

		if index, ok := hashMap; ok {

		}
	}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	res := getTwoSum(nums, target)

	fmt.Println(res)
}
