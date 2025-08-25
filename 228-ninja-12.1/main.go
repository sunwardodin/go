package main

import (
	"fmt"
	"myModule/228-ninja-12.1/dog"
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "fido",
		age:  dog.Years(10),
	}
	fmt.Println(fido)
}
