package main

import "fmt"

func main() {
    target := 5
    integers := []int{1, 2, 3, 4}
    result := twoSum(integers, target)
    fmt.Println("Result:", result)
}

/*
*/
func twoSum(integers []int, target int) []int {
    complCount := make(map[int]int)

    for i, num := range integers {
        complement := target - num
        if complIndex, found := complCount[complement]; found {
            return []int{i, complIndex}
        }
        complement[Index]
    }
    complement[indx]
}
