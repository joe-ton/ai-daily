
package main

import (
    "fmt"
)

type TwoSumInput struct {
    Integers []int `json:"integers"`
    Target int `json:"target"`
}

func main() {
    input := TwoSumInput {
        Integers: []int{1, 2, 3, 4},
        Target: 7,
    }
    i, j, ok := twoSum(input)
    if ok {
        fmt.Printf("Indices: %d and %d\n", i, j)
    } else {
        fmt.Println("Invalid Indices")
    }
}

func twoSum(input TwoSumInput) (int, int, bool) {
    compMap := make(map[int]int)

    for i, num := range input.Integers {
        complement := input.Target - num
        if compIndex, ok := compMap[complement]; ok {
            return compIndex, i, true
        }
        compMap[num] = i
    }
    return 0, 0, false
} 
