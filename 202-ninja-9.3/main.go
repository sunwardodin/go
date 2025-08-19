package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func addTo(i *int) {
	v := i
	*v++
	*i = *v
	wg.Done()
}

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("GoRoutines:", runtime.NumGoroutine())
	incre := 0
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go addTo(&incre)
		fmt.Println("GoRoutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println(incre)
}
