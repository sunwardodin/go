package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	defer fun()

	_, err := os.Open("no-file.txt")
	if err != nil {
		log.Fatalln(err) // fatal will force the program to close once it hits. Nothing will run after it.
	}

	fun()
}

func fun() {
	fmt.Println("Fun run")
}
