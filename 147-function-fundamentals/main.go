package main

import "fmt"

func main() {
	foo()
}

// func (r receiver) identifier(p parameter(s)) (r return(s)) { code }

func foo() {
	fmt.Println("Foo ran")
}
