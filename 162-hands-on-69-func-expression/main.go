package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func main() {
	a := add(1, 2)
	fmt.Println(a)
}
