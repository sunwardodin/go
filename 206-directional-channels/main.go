package main

import "fmt"

func main() {
	c := make(chan<- int, 2)
	cr := make(<-chan int, 2) //receiv
	cs := make(chan<- int, 2) //send

	c <- 42
	cs <- 43

	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", cr)
	fmt.Printf("%T\n", cs)
}
