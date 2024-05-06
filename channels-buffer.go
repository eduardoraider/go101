package main

import (
	"fmt"
	"time"
)

func numbersBuffer(n chan<- int) {
	for i := 1; i <= 10; i++ {
		n <- i
		fmt.Printf("Written in the channel: %d\n", i)
		time.Sleep(time.Millisecond * 100)
	}
	fmt.Println("End of writing")
	close(n)
}

func main() {
	cn := make(chan int, 3)
	go numbersBuffer(cn)

	for v := range cn {
		fmt.Printf("Read from channel: %d\n", v)
		time.Sleep(time.Millisecond * 500)
	}

	fmt.Println("End of execution")

}
