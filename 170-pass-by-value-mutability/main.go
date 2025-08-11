package main

import "fmt"

func intDelta(n *int) {
	*n = 43
}

func main() {
	x := 42
	fmt.Println(x)
	intDelta(&x)
	fmt.Println(x)
}
