package main

import "fmt"

type sumthin struct {
	first string
}

func (s *sumthin) changeFirst(d string) {
	s.first = d
}

func main() {
	dan := sumthin{"Dan"}
	fmt.Println(dan.first)
	dan.changeFirst("Paul")
	fmt.Println(dan.first)
}
