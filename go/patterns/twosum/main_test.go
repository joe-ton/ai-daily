package main

import (
    "errors"
    "testing"
)

func TestTwoSum(t *testing.T) {
    tests := []struct {
        name string
        nums []int
        target int
        wantErr error
        wantResp []int
    } {
        {
            name: "Correct Test",
            nums: []int{1, 2, 3, 4},
            target: 7,
            wantErr: nil,
            wantResp: []int{2, 3},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            gotResp, gotErr := TwoSum(tt.nums, tt.target)
        })
    }

}
