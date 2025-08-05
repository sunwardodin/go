package main

import "fmt"

func foo() int {
	return 27
}

func bar() (int, string) {
	return 42, "Hello"
}

func main() {
	x := foo()
	fmt.Println(x)
	y, f := bar()
	fmt.Println(y, f)
}
