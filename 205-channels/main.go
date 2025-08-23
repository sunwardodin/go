package main

import "fmt"

func main() {
	c := make(chan int)

	// c := make(chan int, 1) this would work with the below command. This tells the channel to accept one value
	// c <- 42 this will not run by itself because the channel has blocked it. You need to command the channel to pull it off

	go func() { // this function actively takes the int off the channel
		c <- 42
	}()

	fmt.Println(<-c)
}
