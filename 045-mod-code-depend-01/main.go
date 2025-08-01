package main

import (
	"fmt"

	puppy "github.com/sunwardodin/puppy1"
)

func main() {
	s1 := puppy.Bark()
	s2 := puppy.Barks()
	fmt.Println(s1)
	fmt.Println(s2)

	fmt.Println(puppy.Bark())
	fmt.Println(puppy.Barks())

}
