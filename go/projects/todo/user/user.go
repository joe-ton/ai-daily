package user

import (
	"bufio"
	"fmt"
	"os"
)

type UserInput struct {
	FirstName string
	LastName  string
	Id        int
}

func AskUserInfo() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("First Name: ")
	firstName, _ := reader.ReadString('\n')

	fmt.Println("Confirmed:", firstName)
}
