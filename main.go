package main

import (
    "fmt"
    "errors"
)

type TwoSumer interface {
    TwoSum(nums []int, target int) ([]int, error)
}

type MapSolver struct {}

func (m MapSolver) TwoSum(nums []int, target int) ([]int, error) {
    seen := make(map[int]int)

    for i, num := range nums {
        complement := target - num
        
        if ta
    }
}


