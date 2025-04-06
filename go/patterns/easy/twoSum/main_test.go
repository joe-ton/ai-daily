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
        wantResp []int
        wantErr error
    } {
        {
            name: "Test1",
            nums: []int{1, 2, 3, 4},
            target: 7,
            wantResp: []int{2, 3},
            wantErr: nil,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            req := TwoSumRequest{nums: []int{}}
            gotResp, gotErr := req.TwoSumRequest()



            
        })
    }
}
