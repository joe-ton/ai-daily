package main

import (
    "errors"
    "fmt"
)

type TwoSumSolver interface {
    TwoSum(nums []int, target int) ([]int, int)
}

type MapSolver struct {}

func TwoSum(nums []int, target int) ([]int, int) {
    seen := make(map[int]int)

    for i, num := range nums {
        complement := target - num

        if 
    }
}
