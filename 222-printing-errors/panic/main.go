package main

import (
	"fmt"
	"os"
)

func main() {
	defer fun()
	_, err := os.Open("no-file.txt")
	if err != nil {
		panic(err)
	}
}

func fun() {
	fmt.Println("Fun run")
}
