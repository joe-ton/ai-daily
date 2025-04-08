package user

import (
	"bufio"
	"fmt"
)

type User struct {
	FirstName string
	LastName  string
	Id        int
}

func AskUser() {
	reader := bufio.NewReader()

	fmt.Println("First Name: ")

}
