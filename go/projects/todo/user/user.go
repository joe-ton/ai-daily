// user/user.go

package user

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

type User struct {
	FirstName string
	LastName  string
	Id        int
}

func (u User) UserAsk() (*User, error) {
    if 

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("First Name: ")
	firstName, err := reader.ReadString('\n')
	if err != nil {
		return nil, errors.New("No First Name")
	}

	fmt.Print("Last Name: ")
	lastName, err := reader.ReadString('\n')
	if err != nil {
		return nil, errors.New("No Last Name")
	}

	return &User{
		FirstName: firstName,
		LastName:  lastName,
		Id:        1,
	}, nil
}
