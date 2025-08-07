package main

import "fmt"

func main() {

	foo()

	func() {
		fmt.Println("Anon func ran")
	}() // These are executing parenthesis. They allow the function to run

	func(s string) {
		fmt.Println("This is my name, ", s)
	}("Drew")
}

func foo() {
	fmt.Println("Foo ran")
}
