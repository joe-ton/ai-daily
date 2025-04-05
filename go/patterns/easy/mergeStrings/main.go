package main

import (
    "fmt"
    "errors"
)

type MergeResponse interface {
    MergeStrings() (string, error)
}

type MergeRequest struct {
    word1 string
    word2 string
}

func (req MergeRequest) MergeResponse() (string, error) {
    resp := "" // output
    // alternate
    // for loop

    for _, char := range req.word1 {
        resp += char
    }
    return resp
}

