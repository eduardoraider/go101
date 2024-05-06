package main

import "fmt"

type Person struct {
	Name   string
	Age    uint8
	Status bool
}

func newPerson(name string) *Person {
	p := Person{Name: name}
	p.Age = 14
	p.Status = true
	return &p

}

func main() {
	var p Person
	p.Name = "Jocko"
	p.Age = 5
	p.Status = true
	fmt.Println(p)

	person := Person{
		Name:   "Malu",
		Age:    10,
		Status: true,
	}
	fmt.Println(person)

	np := newPerson("Gucci")
	fmt.Println(np)
}
