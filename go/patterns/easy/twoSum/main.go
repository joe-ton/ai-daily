package main

import (
    "fmt"
    "errors"
)

type TwoSumInput struct {
    Nums []int
    Target int
}

// Find two indices based on target found from Nums
func TwoSum(input TwoSumInput) ([]int, error) {
    if len(input.Nums) < 2 {
        return nil, errors.New("Invalid count")
    }
    if input.Target < 0 {
        return nil, errors.New("Invalid target")
    }

    seen := make(map[int]int)

    for idx, num := range input.Nums {
        complement := input.Target - num
        if compIdx, found := seen[complement]; found {
            return []int{compIdx, idx}, nil
        }
        seen[num] = idx
    }
    return nil, errors.New("Invalid Indices")
}

func main() {
    input := TwoSumInput{Nums: []int{1, 2, 3, 4}, Target: 7}
    resp, err := TwoSum(input)
    if err != nil {
        fmt.Println("Error:", err.Error())
    } else {
        fmt.Println("Resp:", resp)
    }
}


