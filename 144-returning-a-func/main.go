package main

import "fmt"

func main() {

	x := bar()
	fmt.Println(x())

	fmt.Printf("%T \n", bar())
	fmt.Printf("%T \n", x)

}

func bar() func() int { //This is a function returning a function because functions are just like any other variable
	return func() int {
		return 42
	}
}
