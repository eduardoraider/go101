package main

import (
	"fmt"
	"log"
	"os"
)

var (
	heads, tails int
)

func flipCoin(side string) {
	switch side {
	case "heads":
		heads++
		fmt.Println("heads: ", heads)
	case "tails":
		tails++
		fmt.Println("tails: ", tails)
	default:
		fmt.Println("landed on edge")
	}
}

func main() {

	a, b := 50, 100
	if a < b {
		fmt.Println("a < b")
	} else if a == b {
		fmt.Println("a == b")
	} else if a > b {
		fmt.Println("a > b")
	} else {
		fmt.Println("a > b")
	}

	switch {
	case a < b:
		fmt.Println("a < b")
	case a == b:
		fmt.Println("a == b")
	case a < b:
		fmt.Println("a < b")
	default:
		fmt.Println("a < b")
	}

	flipCoin("heads")

	file, err := os.Open("hello.txt")
	if err != nil {
		log.Panic(err)
	}

	defer file.Close()

	data := make([]byte, 100)
	if _, err := file.Read(data); err != nil {
		log.Panic(err)
	}
	fmt.Println(string(data))

}
