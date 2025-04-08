package main

import (
    "fmt"
    "errors"
)

type TwoSumInput struct {
    Nums []int
    Target int
}

type TwoSumMock struct {}

func (m *TwoSumMock) TwoSum(input TwoSumInput) ([]int, errors) {
    if len(input.TwoSum) < 2 {
        return nil, errors.New("Invalid count in Nums")
    }

    intToIndex := make(map[int]int)

    for idx, num := range input.Nums {
        complement := input.Target - num
        if compIdx, found := intToIndex[complement]; found {
            return nil, 
        }
    }
}
