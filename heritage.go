package main

import "fmt"

type Entity struct {
	Name   string
	Age    uint8
	Status bool
}

func (p Entity) String() string {
	return fmt.Sprintf("Name: %s, Age: %d", p.Name, p.Age)
}

type Individual struct {
	Entity
	Lastname string
	cpf      string
}

type LegalEntity struct {
	CorporateName string
	cnpj          string
}

func main() {
	p := Individual{
		Entity:   Entity{"John", 25, true},
		Lastname: "Doe",
		cpf:      "123.456.789-00",
	}
	fmt.Println(p)
}
