package main

import "fmt"

func main() {
	names := []string{"Eduardo", "Taisa", "Jocko", "Malu"}
	fmt.Println(names)
	fmt.Println(len(names), cap(names))

	names = append(names, "Gucci")
	fmt.Println(names)
	fmt.Println(len(names), cap(names))

	moreNames := make([]string, 3, 4)
	moreNames = append(moreNames, "Eduardo")
	moreNames = append(moreNames, "Jocko")
	moreNames = append(moreNames, "Malu")

	fmt.Println(moreNames)

}
