package main

import "fmt"

func add() int {
	return 42
}

func sum() int {
	return 4 + 2
}

func anotherOne() int {
	return 4 + 4
}

func main() {
	defer fmt.Println(add())
	defer fmt.Println(anotherOne())
	fmt.Println(sum())
}
