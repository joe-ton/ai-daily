package main

import (
    "fmt"
    "errors"
)

type TwoSumSolver interface {
    TwoSum(nums []int, target int) ([]int, error)
}

type MapSolver struct {}

func (m MapSolver) TwoSum(nums []int, target int) ([]int, error) {
    seen := make(map[int]int)

    for i, num := range nums {
        complement := num - i
        if compIdx, found := seen[complement]; found {
            return []int{compIdx, i}, nil
        }
        seen[num] = i
    }
    return nil, errors.New("Invalid")
}

func main() {
    solver := MapSolver{}
    resp, err := solver.TwoSum([]int{1, 2, 3, 4}, 7)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Resp:", resp)
    }
}
