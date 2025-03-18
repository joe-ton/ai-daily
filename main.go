package main

import "fmt"

func main() {
	integers := []int{2, 7, 11, 15}
	target := 9
	result := getTwoSum(target, integers)
	fmt.Println("Indices:", result)
}

/*
Find two integers from slice to add up to targeted sum.

Parameters:
- target int: total sum needed to add up to
- integers []int: integers given

Returns:
- []int: empty or 2 slices of indices whose values add to targeted sum

Tradeoff Complexities:
1. Brute Force: Time O(n^2), Space O(1)
-> 2. HashMap: Time O(n), Space O(n)
3. Sorted / 2 pointer: Time O(nlogn), Space O(n) or O(1)
*/
func getTwoSum(target int, integers []int) []int {
	numMap := make(map[int]int) // complement counter

	for i, num in range integers {
		complement := target - num
		// found { } is called a comma-ok idiom, generates a boolean
		if compIndex, found := numMap[complement]; found {
			return []int{compIndex, i}
		}
		numMap[num] = i
	}
	return []int{}
}




