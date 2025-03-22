package main

import "fmt"

func main() {
    integers := []int{1, 2, 3, 4}
    target := 5
    result := twoSum(integers, target)
    fmt.Println("Result:", result)
}

/*
Find two values from given slice adding to target sum.

Parameters:
- integers []int: given integers
- target int: target sum

Returns:
- []int: indices

Solutions:
1. Brute Force: Time (n^2), Space (1)
2. Map: Time (n), Space (n)
3. Sorted: Time (nlogn), Space(n)
*/
func twoSum(integers []int, target int) []int {
    numMap := make(map[int]int)

    for i, num := range integers {
        complement := target - num
        if complIdx, found := numMap[complement]; found {
            return []int{i, complIdx}
        }
        numMap[num] = i
    }

    return []int{}
}
