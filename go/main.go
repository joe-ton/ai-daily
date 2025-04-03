package main

import (
    "fmt"
    "errors"
)

type MapSolver struct {}

type TwoSumInterface interface {
    TwoSum(nums []int, target int) ([]int, error)
}

func (m MapSolver) TwoSum(nums []int, target int) ([]int, error) {
    seen := make(map[int]int)

    for i, num := range nums {
        complement := target - num

        if compIdx, found := seen[complement]; found {
            return []int{compIdx, i}, nil
        }
        seen[num] = i
    }
    return nil, errors.New("Invalid Number")
}

func main() {
    solver := TwoSumInterface(MapSolver{})
    nums := []int{1, 2, 3, 4}
    target := 7

    resp, err := solver.TwoSum(nums, target)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Response:", resp)
}
