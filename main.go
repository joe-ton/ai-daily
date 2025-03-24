package main

import (
    "fmt"
)

type twoSumData struct {
    Integers []int `json:"integers"`
    Target int `json:"target"`
}

func main() {
    input := twoSumData{
        Integers: []int{1, 2, 3, 4},
        Target: 9,
    }
}
