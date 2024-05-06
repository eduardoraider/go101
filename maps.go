package main

import (
	"fmt"
	"maps"
)

func main() {
	m := make(map[string]int)

	m["a"] = 1
	m["b"] = 2
	m["c"] = 3

	value, exists := m["a"]
	println(value, exists)

	v, ok := m["d"]
	println(v, ok)

	fmt.Println(m)
	fmt.Println(m["b"])
	fmt.Println(len(m))

	delete(m, "a")
	fmt.Println(m)
	fmt.Println(len(m))

	n := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println("map:", n)

	n2 := map[string]int{"a": 1, "b": 2, "c": 3}
	if maps.Equal(n, n2) {
		fmt.Println("true")
	}
}
