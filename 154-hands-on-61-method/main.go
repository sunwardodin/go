package main

import "fmt"

type person struct {
	first string
	age   int
}

func (p person) Speak() {
	fmt.Println("My name is", p.first, "and my age is", p.age)
}

func main() {
	p := person{
		first: "George",
		age:   55,
	}

	p.Speak()
}
