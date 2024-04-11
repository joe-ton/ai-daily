package main

import (
	"errors"
	"fmt"
)

func addTwoIntegers(a, b int) (int, error) {
	if a < 0 {
		return 0, errors.New("Hello, wrong")
	}
	return a + b, nil
}

func main() {
	result, err := addTwoIntegers(5, 8)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(result)
}
