package main

import (
	"fmt"
	"time"
)

func numbers() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d ", i)
		time.Sleep(time.Millisecond * 150)
	}
}

func letters() {
	for l := 'a'; l <= 'j'; l++ {
		fmt.Printf("%c ", l)
		time.Sleep(time.Millisecond * 230)
	}
}

func main() {
	go numbers()
	go letters()
	time.Sleep(1 * time.Second)
	fmt.Println("End of execution")

}
