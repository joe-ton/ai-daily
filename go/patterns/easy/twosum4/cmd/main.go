package main

import (
	"fmt"

	"github.com/joe-ton/solution"
)

func main() {
	nums := []int{1, 2, 3, 4}
	target := 7
	twoSum := solution.TwoSum{Nums: &nums, Target: &target}
	resp, err := twoSum.Find()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Response:", resp)
	}
}
