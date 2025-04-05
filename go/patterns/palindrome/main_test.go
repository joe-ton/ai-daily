package main

import (
    "testing"
)

func TestIsPalindrome(t *testing.T) {
    tests := []struct {
        name string
        s string
        wantBool bool
    } {
        {
            name: "Test1",
            s: "racecar",
            wantBool: true,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := IsPalindrome(tt.)
        })
    }
}
