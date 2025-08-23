package main

import "fmt"

func main() {
	c := make(chan int)

	// send
	go send(c)
	// receive
	go receive(c)

	fmt.Print("About to exit\n")
}

func send(c chan<- int) {
	c <- 42
}

func receive(c <-chan int) {
	fmt.Println(<-c)
}
