package main

import (
    "testing"
    "errors"
)

func TestIsPalindrome(t *testing.T) {
    tests := []struct {
        name string
        s string
        wantError error
        wantResp bool
    } {
        {
            name: "Test1",
            s: "racecar",
            wantError: nil,
            wantResp: true,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {

        })
    }
}
