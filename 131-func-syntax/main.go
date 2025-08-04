package main

import "fmt"

// func (r receiver) identifier(p parameter(s)) (return(s)) { <code> }

func main() {
	foo()

	bar("Drew")

	s1 := aloha("George")
	fmt.Println(s1)

	fmt.Println(dogYears("Drew", 29))
}

func foo() {
	fmt.Println("I am from foo")
}

func bar(s string) {
	fmt.Println("My name is", s)
}

func aloha(s string) string {
	return fmt.Sprint("My name is ", s)
}

func dogYears(name string, age int) (string, int) {
	age *= 7
	return fmt.Sprint(name, " is this old in dog years:"), age
}
