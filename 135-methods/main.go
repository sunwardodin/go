package main

import "fmt"

type person struct {
	first string
}

func (p person) speak() { //the p is the identifier for the type. It can actually be anything you want
	fmt.Println("I am", p.first) // notice how the p is used here in p.first
}

func main() {
	p1 := person{
		first: "James",
	}
	p2 := person{
		first: "Jenny",
	}

	p1.speak()
	p2.speak()
}
