package main

import (
    "fmt"
    "errors"
)

type TwoSumSolver interface {
    Run() ([]int, error)
}

type TwoSumData struct {
    Nums []int
    Target int
}

func NewTwoSumData(nums []int, target int) *TwoSumData {
    return &TwoSumData{
        Nums: nums,
        Target: target,
    }
}

func (d TwoSumData) TwoSumBrute() ([]int, error) {
    if len(d.Nums) < 2 {
        return nil, errors.New("Invalid count")
    }

    for i := 0; i < len(d.Nums); i++ {
        for j := 0; j < len(d.Nums)-1; j++ {
            return []int{d.Nums[i], d.Nums[j]}, nil
        }
    }
    return nil, errors.New("Invalid Indices")
}

func (d TwoSumData) TwoSumHash() ([]int, error) {
    if len(d.Nums) < 2 {
        return nil, errors.New("Invalid count")
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

func (d *TwoSumData) Run() ([]int, error) {
    return d.TwoSumHash()
}

func main()  {
    data := TwoSumData{Nums: []int{1, 2, 3, 4}, Target: 7}
    resp, err := data.Run()

    if err != nil {
        fmt.Println("Error:", err.Error())
    } else {
        fmt.Println("Response:", resp)
    }
}
