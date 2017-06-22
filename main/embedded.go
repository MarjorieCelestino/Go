package main

import "fmt"

type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

type Android struct {
	// é tipo "herança"
	Person
	Model string
}

func main() {
	a := new(Android)

	// Pode acessar o método "herdado" assim
	a.Person.Talk()

	// Ou assim
	a.Talk()
}
