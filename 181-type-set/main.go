package main

import "fmt"

func addInt(a int, b int) int {
	return a + b
}

func addFloat(a float64, b float64) float64 {
	return a + b
}

type myNumbers interface { // this becomes a type set that you can input into a function in place of the regular pipe
	int | float64
}

func addT[T myNumbers](a, b T) T {
	return a + b
}

func main() {
	fmt.Println(addInt(2, 3))
	fmt.Println(addFloat(6.2, 5.3))

	fmt.Println(addT(2, 3))
	fmt.Println(addT(6.2, 5.3))

}
