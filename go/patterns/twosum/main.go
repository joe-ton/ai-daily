package main

import (
    "errors"
    "fmt"
)

type TwoSumSolver interface {
    TwoSum(nums []int, target int) ([]int, error)
}

type MapSolver struct {}


func (m MapSolver) TwoSum(nums []int, target int) ([]int, error) {
    if len(nums) < 2 {
        return nil, errors.New("Invalid number of values")
    }
    if target < 0 {
        return nil, errors.New("Invalid target sum")
    }

    seen := make(map[int]int)

    for idx, num := range nums {
        complement := target - num
        if compIdx, found := seen[complement]; found {
            return []int{compIdx, idx}, nil
        }
        seen[num] = idx
    }
    return nil, errors.New("Invalid Indices")
}

func main() {
    nums := []int{1, 2, 3, 4}
    target := 7

    solver := MapSolver {}
    resp, err := solver.TwoSum(nums, target)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Response:", resp)
}
