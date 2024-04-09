package main

import (
	"fmt"

	"github.com/joe-ton/goLC/practice"
)

func main() {
	nums := []int{1, 2, 3, 1}
	target := 5
	result := practice.GetTwoSum(nums, target)
	fmt.Println(result)
}
