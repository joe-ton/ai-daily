package main

import (
    "fmt"
    "errors"
)

type TwoSumSolver interface {
    TwoSum() ([]int, error)
}

type TwoSumData struct {
    Nums []int
    Target int
}

func (d TwoSumData) TwoSumSolver() ([]int, error) {
    if len(d.Nums) < 2 {
        err := errors.New("Not enough values in the Nums")
        return nil, err
    }

    seen := make(map[int]int)

    for idx, num := range d.Nums {
        complement := d.Target - num
        if compIdx, found := seen[complement]; found {
            return []int{compIdx, idx}, nil
        }
        seen[num] = idx
    }
    return nil, errors.New("Invalid Indices")
}

func main() {
    data := TwoSumData{Nums: []int{1, 2, 3, 4}, Target: 7}
    resp, err := data.TwoSumSolver()
    if err != nil {
        fmt.Println("Error:", err.Error())
    } 
    fmt.Println("Response:", resp)
}
