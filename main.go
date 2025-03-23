package main

import (
    "fmt"
    "net/http"
)

func main() {
    resp, err := http.Get("http://www.google.com")
    fmt.Println("Resp:", resp)
    fmt.Println("Err:", err)
}
