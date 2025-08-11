package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	fmt.Println(&x) // The & gives you the address of where x is stored in memory
}
