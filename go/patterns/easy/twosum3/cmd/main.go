package main

import (
	"fmt"

	"github.com/joe-ton/solution"
)

func main() {
	s := solution.TwoSumSolver{Nums: []int{1, 2, 3, 4}, Target: 7}
	resp, err := s.Solve()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Response:", resp)
	}
}
