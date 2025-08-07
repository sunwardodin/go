package main

import "fmt"

func main() {
	x := func() {
		fmt.Println("Anon func ran")
	}

	x()

	y := func(s string) {
		fmt.Println("This is my name, ", s)
	}

	y("Drew")
}
