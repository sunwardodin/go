package main

import "fmt"

func main() {
	c := make(chan int)

	// send
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c) // this is how you close a channel
	}()

	// receive
	for v := range c { // range will keep trying to pull values off of a channel until it is closed
		fmt.Println(v)
	}

	fmt.Print("About to exit\n")
}
