package main

import "fmt"

func main() {
	fmt.Println(printSquare(square, 2))

}

func square(n int) int {
	return n * n
}

func printSquare(f func(int) int, a int) string {
	x := f(a)
	return fmt.Sprintf("This number %v squared is %v", a, x)
}
