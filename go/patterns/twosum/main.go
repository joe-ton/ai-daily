package main

import (
    "fmt"
    "errors"
    "log"
    "time"

    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
    twoSumCalls = prometheus.NewCounter(

    )
)

type TwoSumSolver interface {
    TwoSum(nums []int, target int) ([]int, error)
}

type MapSolver struct {}


func (m MapSolver) TwoSum(nums []int, target int) ([]int, error) {
    log.Printf("TwoSum called with num=%v and target=%d", nums, target)
    if len(nums) < 2 {
        err := errors.New("Invalid Nums: not enough values in nums")

        log.Println("Error:", err)
        return nil, err
    }


    seen := make(map[int]int) // num -> idx lookup

    for idx, num := range nums {
        complement := target - num 
        if compIdx, found := seen[complement]; found {
            return []int{compIdx, idx}, nil
        }
        seen[num] = idx
    }
    return nil, errors.New("Invalid indices")
}

func main() {
    solver := MapSolver {}
    resp, err := solver.TwoSum([]int{1, 2, 3, 4}, 7)
    if err != nil {
        fmt.Println("Error:", err.Error())
    }
    fmt.Println("Response:", resp)
}
