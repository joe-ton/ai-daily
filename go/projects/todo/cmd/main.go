// cmd/main.go

package main

import (
	"fmt"

	"github.com/joe-ton/ai-daily/user"
)

func main() {
	person, err := user.UserAsk()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

}
