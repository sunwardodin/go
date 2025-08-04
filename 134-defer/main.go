package main

import "fmt"

func main() {
	defer foo() // defer tells the function "don't run this until the end of this function"
	bar()
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}
