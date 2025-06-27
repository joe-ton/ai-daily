package solution

type Greeter interface {
	Greet() (string, error)
}

type Person struct {
	Name string
}

func (p *Person) Greet() (string, error) {
	greeting := "Hello, " + p.Name
	return greeting, nil
}
