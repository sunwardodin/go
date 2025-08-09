package main

import "fmt"

func first() func() int {
	return func() int {
		return 42
	}
}

func main() {
	a := first()
	fmt.Println(a())
}
