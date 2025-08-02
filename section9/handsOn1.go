package main

import "fmt"

func Bark() string {
	s1 := "Woof!"
	return s1
}

func main() {
	fmt.Printf("Here is: %v \n", Bark())
}
