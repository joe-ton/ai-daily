package main

import (
    "fmt"
    "errors"
)

type TwoSumInput struct {
    Nums []int
    Target int
}

type BruteSolver struct{}

func (b *BruteSolver) Solve(input TwoSumInput) ([]int, error) {
    fmt.Println("Testing1")
    if len(input.Nums) < 2 {
        return nil, errors.New("Invalid count of nums")
    }

    for i := 0; i < len(input.Nums); i++ {
        for j := i + 1; j < len(input.Nums); j++ {
            total := input.Nums[i] + input.Nums[j]
            if total == input.Target {
                return []int{i, j}, nil
            }
        }
    }
    return nil, errors.New("Invalid indices")
}

type MapSolver struct {}

// Find two indices based of nums to sum to target
func (m *MapSolver) Solve(input TwoSumInput) ([]int, error) {
    if len(input.Nums) < 2 {
        return nil, errors.New("Invalid count of nums")
    }
    intToIdx := make(map[int]int)

    for idx, num := range input.Nums {
        complement := input.Target - num
        if compIdx, found := intToIdx[complement]; found {
            return []int{compIdx, idx}, nil
        }
        intToIdx[num] = idx
    }
    return nil, errors.New("Invalid Indices")
}

func main() {
    input := TwoSumInput{Nums: []int{1, 2, 3, 4}, Target: 7}

    solver := &MapSolver{}
    resp, err := solver.Solve(input)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Response:", resp)
    }
}
