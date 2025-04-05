package main

import (
	"errors"
	"fmt"
)

type TwoSumResponse interface {
    TwoSum() ([]int, error) // indices
}

type TwoSumRequest struct {
    nums []int
    target int
}

// return indices from nums
func (req TwoSumRequest) TwoSum() ([]int, error) {
    if len(req.nums) < 2 {
        return nil, errors.New("Invalid elements of nums")
    }
    seen := make(map[int]int)

    for idx, num := range req.nums {
        complement := req.target - num

        if compIdx, found := seen[complement]; found {
            return []int{compIdx, idx}, nil
        }
        seen[num] = idx
    }
    return nil, errors.New("Invalid Indices")
}

func main() {
    req := TwoSumRequest{nums: []int{1, 2, 3, 4}, target: 7}
    resp, err := req.TwoSum() 
    if err != nil {
        fmt.Println("Error:", err.Error())
    }
    fmt.Println("Response:", resp)
}
