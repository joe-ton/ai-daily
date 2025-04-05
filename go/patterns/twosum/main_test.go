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
        wantErrors error
        wantResp []int
    } {
        {
            name: "Test1",
            nums: []int{1, 2, 3, 4},
            target: 7,
            wantErrors: nil,
            wantResp: []int{2, 3},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {

        })
    }
}
