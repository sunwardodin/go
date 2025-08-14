package main

import "fmt"

func addInt(a int, b int) int {
	return a + b
}

func addFloat(a float64, b float64) float64 {
	return a + b
}

func addT[T int | float64](a, b T) T { // this is a generic type function. The brackets allow you to specify that this function will receive either type into the function.
	return a + b
}

func main() {
	fmt.Println(addInt(2, 3))
	fmt.Println(addFloat(6.2, 5.3))

	fmt.Println(addT(2, 3))
	fmt.Println(addT(6.2, 5.3))

}
