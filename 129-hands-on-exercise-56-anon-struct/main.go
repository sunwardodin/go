package main

import "fmt"

func main() {
	p1 := struct {
		name string
		hair string
		age  int
	}{
		name: "George",
		hair: "grey",
		age:  26,
	}

	fmt.Printf("%T \n", p1)
}
