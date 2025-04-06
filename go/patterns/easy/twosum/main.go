package main

import (
    "fmt"
    "errors"
)

type TwoSumData struct {
    Nums []int
    Target int
}

func TwoSum(data TwoSumData) ([]int, error) {
    if len(data.Nums) < 2 {
        return nil, errors.New("Invalid count")
    }
    seen := make(map[int]int)

    for idx, num := range data.Nums {
        complement := data.Target - num
        if compIdx, found := seen[complement]; found {
            return []int{compIdx, idx}, nil // early return
        }
        seen[num] = idx // tracker
    }
    return nil, errors.New("Invalid indices") // fallback
}

func main() {
    data := TwoSumData{Nums: []int{1, 2, 3, 4}, Target: 7}
    resp, err := TwoSum(data)
    if err != nil {
        fmt.Println("Error:", err.Error())
    } else {
        fmt.Println("Response:", resp)
    }
}

