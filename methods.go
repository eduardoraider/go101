package main

import "fmt"

type Category struct {
	Name   string
	Parent *Category
}

type User struct {
	Name   string
	Age    uint8
	Status bool
	cpf    string
}

func (u User) String() string {
	return fmt.Sprintf("Name: %s, Age: %d, Status: %v", u.Name, u.Age, u.Status)
}

// Methods are linked to structs

func (c Category) HasParent() bool {
	return c.Parent != nil
}

func (c *Category) SetParent(parent *Category) {
	c.Parent = parent
}

func main() {

	u := User{"John", 25, true, "123.456.789-00"}
	fmt.Println(u)

	cat := Category{Name: "category 1"}
	if !cat.HasParent() {
		fmt.Println("Category 1 doesn't have parent")
	}

	cat2 := Category{Name: "category 2"}
	cat2.SetParent(&Category{Name: "category 3"})
	if !cat2.HasParent() {
		fmt.Println("Category 2 doesn't have parent")
	}

}
