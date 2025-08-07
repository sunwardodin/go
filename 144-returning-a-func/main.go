/*
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
*/

package main

import "fmt"

// Passing 'add' function as argument
// This function returns another function
func incrementer() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}
func main() {
	increment := incrementer()
	fmt.Println(increment()) // Output: 1
	fmt.Println(increment()) // Output: 2
	fmt.Println(increment()) // Output: 3
}
