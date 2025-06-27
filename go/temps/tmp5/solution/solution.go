package solution

import "errors"

type Greeter interface {
	Greet(phrase string) (string, error)
}

type Person struct {
	PersonName string
}

var _ Greeter = (*Person)(nil)

func (p *Person) Greet(phrase string) (string, error) {
	if phrase == "" {
		return "", errors.New("no phrase")
	}
	return phrase + " " + p.PersonName, nil
}
