package main

import (
	"fmt"
)

/*
Find two integers from slice to add to sum of target.

Parameters:
- 
*/
func twoSum(integers []int, target int) []int {
	numMap := make(map[int]int)
	for i, num := range integers {
		complement := target - num
		indexCompl, found := numMap[complement]; found {
			[]int{num, indexCompl}
		}
	}
}

func main() {
	integers := []int{2, 7, 11, 15}
	result := twoSum(integers, 9)
	fmt.Println(result)
}
