package main

import (
	"errors"
)

type TwoSumInput struct {
	Nums   []int
	Target int
}

func (i TwoSumInput) TwoSum() ([]int, error) {
	if len(i.Nums) < 2 {
		return nil, errors.New("")
	}
}
