package main

import (
	"fmt"
	"strconv"
)

func hello(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

func sum(a, b int) int {
	return a + b
}

func convertAndSum(a int, b string) (total int, err error) {
	i, err := strconv.Atoi(b)
	if err != nil {
		return
	}

	total = sum(a, i)

	return
}

func main() {
	hello("Eduardo")

	fmt.Println("Total:", sum(1, 2))
	total, err := convertAndSum(1, "10")
	fmt.Println(total, err)
}
