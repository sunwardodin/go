package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	n, err := fmt.Println("Hello") // this will print Hello and the number 6 because of the new line
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)

	var answer1, answer2, answer3 string

	fmt.Print("Name: ")
	_, err = fmt.Scan(&answer1)
	if err != nil {
		panic(err)
	}

	fmt.Print("Fav Food: ")
	_, err = fmt.Scan(&answer2)
	if err != nil {
		panic(err)
	}

	fmt.Print("Fav Sport: ")
	_, err = fmt.Scan(&answer3)
	if err != nil {
		panic(err)
	}

	f, err := os.Create("names.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	r := strings.NewReader("James Bond")
	io.Copy(f, r) // this will read from r and put it into f
}
