package main

import "fmt"

type dog struct {
	first string
}

func (d dog) walk() {
	fmt.Printf("I am %v and I am walking\n", d.first)
}

func (d dog) run() {
	fmt.Println("I am", d.first, "and I am running")
}

type youngin interface {
	run()
	walk()
}

func youngRun(y youngin) {
	y.walk()
	y.run()
}

func main() {
	d1 := dog{"Fido"}
	youngRun(d1)

	d2 := dog{"Bill"}
	youngRun(d2)
	d2 = d2.changeName("rover")
	youngRun(d2)
}

func (d dog) changeName(s string) dog {
	d.first = s
	return d

}
