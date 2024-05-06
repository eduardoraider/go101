package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("========")

	names := []string{"Eduardo", "Taisa", "Jocko", "Malu"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	fmt.Println("========")

	for K, name := range names {
		fmt.Println(K, name)
	}

	fmt.Println("========")

	for _, name := range names {
		fmt.Println(name)
	}

	fmt.Println("========")

	var i int
	for i < len(names) {
		fmt.Println(names[i])
		i++
	}

}
