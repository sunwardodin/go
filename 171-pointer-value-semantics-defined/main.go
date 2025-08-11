package main

import "fmt"

func addOne(i int) int {
	i += 1
	return i
}

func addOneP(i *int) {
	*i += 1
}

func main() {

	// value semantics
	a := 1
	fmt.Println("Value semantics")
	fmt.Println(a)
	fmt.Println(addOne(a))
	fmt.Println(a)

	// pointer semantics
	b := 1
	fmt.Println("\nPointer semantics")
	fmt.Println(b)
	addOneP(&b)
	fmt.Println(b)
}
