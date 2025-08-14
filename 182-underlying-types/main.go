package main

import "fmt"

func addInt(a int, b int) int {
	return a + b
}

func addFloat(a float64, b float64) float64 {
	return a + b
}

type myNumbers interface {
	~int | ~float64
}

func addT[T myNumbers](a, b T) T {
	return a + b
}

type myAlias int // this allows you to double up on types

func main() {
	var n myAlias = 42

	fmt.Println(addInt(2, 3))
	fmt.Println(addFloat(6.2, 5.3))

	fmt.Println(addT(n, 3))
	fmt.Println(addT(6.2, 5.3))

}
