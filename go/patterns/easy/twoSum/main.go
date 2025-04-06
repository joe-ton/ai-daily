package main

import (
    "fmt"
    "errors"
)

type TwoSumSolver interface {
    TwoSum() ([]int, error)
}

type TwoSumRequest struct {
    Nums []int
    Target int
}

func (req TwoSumRequest) TwoSum() ([]int, error) {
    if len(req.Nums) < 2 {
        return nil, errors.New("Invalid count in nums")
    }
    seen := make(map[int]int) // lookup for complements

    for idx, num := range req.Nums {
        complement := req.Target - num
        if compIdx, found := seen[complement]; found {
            return []int{compIdx, idx}, nil
        }
        seen[num] = idx
    }
    return nil, errors.New("Invalid Indices")
}

func main() {
    req := TwoSumRequest{Nums: []int{1, 2, 3, 4}, Target: 7}
    resp, err := req.TwoSum()
    if err != nil {
        fmt.Println("Error:", err.Error())
    }
    fmt.Println("Response:", resp)

}
