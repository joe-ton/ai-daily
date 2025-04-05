package main

import (
    "fmt"
)

func isPalindrome(s string) bool {
    runes := []rune(s)
    n := len(runes)

    for i := 0; i < n/2; i++ {
        if s[i] != s[n-1-i] {
            return false
        }
    }
    return true
}

func main() {
    s := "racecarr"
    resp := isPalindrome(s)
    fmt.Println("Response:", resp)
}
