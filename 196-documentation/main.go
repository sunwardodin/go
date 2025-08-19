package main

import (
	"fmt"
)

func doSomething(x int) int {
	return x * 5
}

func main() {
	ch := make(chan int)
	fmt.Printf("%T\n", ch)
	go func() {
		ch <- doSomething(5)
	}()
	fmt.Println(<-ch)
}
