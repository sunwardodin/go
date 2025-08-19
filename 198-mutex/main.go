package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("GoRoutines:", runtime.NumGoroutine())

	counter := 0

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	var mu sync.Mutex // This is the Mutually Exclusive variable

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock() // This will lock the function into finishing
			v := counter
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock() // This will unlock the function so that the next one can run
			wg.Done()
		}()
		fmt.Println("GoRoutines:", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("GoRoutines:", runtime.NumGoroutine())
	fmt.Println("Counter:", counter)
}
