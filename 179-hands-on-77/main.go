package main

import "fmt"

type person struct {
	first string
}

func (p *person) changeFirst(d string) {
	p.first = d
}

func changeName(p person, s string) person {
	p.first = s
	return p
}

func main() {
	dan := person{"Dan"}
	fmt.Println(dan.first)
	dan.changeFirst("Paul")
	fmt.Println(dan.first)

	dan = changeName(dan, "Jim")
	fmt.Println(dan)
}
