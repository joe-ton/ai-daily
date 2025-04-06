package main

import (
    "fmt"
    "errors"
)

type TwoSumInput struct {
    Nums []int
    Target int
}

func TwoSum(input TwoSumInput) ([]int, error) {
    if len(input.Nums) < 2 {
        return nil, errors.New("Invalid count")
    }
    seen := make(map[int]int) 

    for idx, num := range input.Nums {
        complement := input.Target - num
        if compIdx, found := seen[complement]; found {
            return []int{compIdx, idx}, nil // early 
        }
        seen[num] = idx // tracker
    }
    return nil, errors.New("Invalid indices")
}

func main() {
    input := TwoSumInput{Nums: []int{1, 2, 3, 4}, Target: 7}
    resp, err := TwoSum(input)
    if err != nil {
        fmt.Println("Error:", err.Error())
    } else {
        fmt.Println("Response:", resp)
    }

}
