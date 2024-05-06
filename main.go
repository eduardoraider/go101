package main

import "fmt"

var name string

var (
	username string
	n1       int
	n2       int32
)

func main() {
	name = "Eduardo Raider"
	usename := "eduraider"
	fmt.Println("Hello,", name, "!")

	var a, b, c = true, 2.3, "hello"
	fmt.Println(a, b, c)

	fmt.Println("Hello,", usename, "!")

	var total int
	total++
	fmt.Println("Total", total)

	var x, y = 1, 2
	fmt.Println(x, y)
	x, y = y, x
	fmt.Println(x, y)
}
