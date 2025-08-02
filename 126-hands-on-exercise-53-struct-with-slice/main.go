package main

import "fmt"

type person struct {
	first string
	last  string
	favIC []string
}

func main() {

	p1 := person{
		first: "James",
		last:  "Bond",
		favIC: []string{"chocolate", "banana", "cotton candy"},
	}

	p2 := person{
		first: "George",
		last:  "Bond",
		favIC: []string{"chocolate", "banana", "cotton candy"},
	}

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p1.favIC)
	fmt.Println(p2.favIC)

	for _, v := range p1.favIC {
		fmt.Println(p1.first, " favorite is ", v)
	}

	for _, v := range p2.favIC {
		fmt.Println(p2.first, " favorite is ", v)
	}

}
