package main

import (
    "fmt"
    "errors"
)

func TwoSum(nums []int, target int) ([]int, error) {
    if len(nums) < 2 {
        return nil, errors.New("Not enough content in nums")
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

    resp, err := TwoSum(nums, target)
    if err != nil {
        fmt.Println("Error:", err.Error())
    }
    fmt.Println("Response:", resp)

}
