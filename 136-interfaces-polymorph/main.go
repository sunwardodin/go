package main

import "fmt"

type person struct {
	first string
}

type secretAgent struct {
	person
	ltk bool
}

type farmer struct {
	person
	tools int
}

func (p person) speak() {
	fmt.Println("I am", p.first)
}

func (sa secretAgent) speak() {
	fmt.Println("I am a secret agent", sa.first)
}

func (f farmer) accent() {
	fmt.Println("Y'ALL, ma name is", f.first)
}

type human interface {
	speak()
}

type accent interface { // <--
	accent()
}

func saySomething(h human) {
	h.speak()
}

func downSouth(a accent, h human) { // I made another type interface to see if I could double up on types and it worked
	a.accent()
	h.speak()
}

func main() {
	sa1 := secretAgent{
		person: person{first: "James"},
		ltk:    true,
	}

	p2 := person{
		first: "Jenny",
	}

	f1 := farmer{
		person: person{first: "Jimmy"},
		tools:  4,
	}

	//sa1.speak()
	//p2.speak()

	saySomething(sa1)
	saySomething(p2)
	downSouth(f1, f1)
}
