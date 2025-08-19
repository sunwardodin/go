package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(2) // This tells the wait group how many functions to wait for before ending

	go func() {
		fmt.Println("One ran")
		wg.Done()
	}()
	go func() {
		fmt.Println("two ran")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Main ran")
}
