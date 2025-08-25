package main

import (
	"errors"
	"fmt"
)

var errFoo = errors.New("clearly you are mistaken")

func main() {
	var answer string
	fmt.Print("What's your name? ")
	fmt.Scanln(&answer)

	v := foo(answer)
	if !v {
		fmt.Println(errFoo)
	} else {
		fmt.Println("Ok")
	}
}

func foo(a string) bool {
	return a == "Drew"
}
