package main

import (
    "fmt"
    "errors"
)

type PalindromeSolver interface {
    IsPalindrome() (bool, error)
}

type PalindromeData struct {
    S string
}

func (d PalindromeData) IsPalindrome() (bool, error) {
    runes := []rune(d.S)
    n := len(runes)

    for i := 0; i < n/2; i++ {
        if d.S[i] != d.S[n-1-i] {
            return false, errors.New("Comparison is wrong")
        }
    }
    return true, nil
}

func main() {
    data := PalindromeData{S: "racecar"}
    resp, err := data.IsPalindrome()
    if err != nil {
        fmt.Println("Error:", err.Error())
    }
    fmt.Println("Response:", resp)

}
