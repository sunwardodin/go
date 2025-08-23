package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		c <- 42
		close(c) // this is not strictly necessary in this code
	}()
	v, ok := <-c

	fmt.Println(v, ok)
	fmt.Println(<-c)
	fmt.Println(c)
}
