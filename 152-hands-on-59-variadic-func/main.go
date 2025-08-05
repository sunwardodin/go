package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	y := bar(1, 2, 3, 4, 5, 6, 7, 8)

	fmt.Println(foo(x...))
	fmt.Println(y)
}

func foo(a ...int) int {
	n := 0
	for _, v := range a {
		n += v
	}
	return n
}

func bar(a ...int) int {
	n := 0
	for _, v := range a {
		n += v
	}
	return n
}
