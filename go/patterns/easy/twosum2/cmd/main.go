package main

import (
	"fmt"

	"github.com/joe-ton/solution"
)

func main() {
	input := solution.TwoSumInput{Nums: []int{1, 2, 3, 4}, Target: 7}
	resp, err := input.FindTwoSum()
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Response:", resp)
}
