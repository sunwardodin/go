package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	fmt.Println(&x)

	y := &x
	fmt.Printf("%v \t %T\n", y, y)
	fmt.Println(*y)  // This will show you the value of x not *y
	fmt.Println(*&x) // This dereferences the address of x to show you the value of x instead of the address

	*y = 43 // This will directly change the value of x since y is a pointer to x
	fmt.Println(x)
	fmt.Println(*y)
}
