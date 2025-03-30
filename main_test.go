package main

import (
    "testing"
    "errors"
)

func TestTwoSum(t *testing.T) {
    tests := []struct {
        name string
        nums []int
        target int
        wantIndices []int
        wantError error
    } {
        {
            name: "Testing",
            nums: []int{1, 2, 3, 4},
            target: 7,
            wantIndices: []int{2, 3},
            wantError: nil,
        },
    }
    for _, tt := range tests {
        t.Run()
    }
    
}
