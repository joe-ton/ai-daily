package main

import (
	"errors"
)

type TwoSumInput struct {
	Nums   []int
	Target int
}

type TwoSumStruct struct{}

// Find two indices whose values from input to equal target
func TwoSum(input TwoSumInput) ([]int, error) {
	if len(input.Nums) < 2 {
		return nil, errors.New("Invalid count on nums")
	}

}
