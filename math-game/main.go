package main

import (
	"fmt"
	"math/rand"
)

// Math Game
func genProb() bool {
	sl := []string{"+", "-", "/", "*"}
	randomOp := sl[rand.Intn(3)]

	n1 := rand.Intn(100)
	n2 := rand.Intn(100)
	answer := 0
	cAnswer := 0

	if randomOp == "+" {
		cAnswer = n1 + n2
		fmt.Printf("What is %v + %v\n", n1, n2)
		fmt.Scanln(&answer)
		if answer != cAnswer {
			return false
		}
	} else if randomOp == "-" {
		cAnswer = n1 - n2
		fmt.Printf("What is %v - %v\n", n1, n2)
		fmt.Scanln(&answer)
		if answer != cAnswer {
			return false
		}
	} else if randomOp == "*" {
		cAnswer = n1 * n2
		fmt.Printf("What is %v * %v\n", n1, n2)
		fmt.Scanln(&answer)
		if answer != cAnswer {
			return false
		}
	}

	return true
}

func main() {
	var a string
	var b bool
	i := 0

	fmt.Print("Ready for a problem?: ")
	fmt.Scanln(&a)
	if a == "yes" || a == "Yes" {
		for i <= 1 {
			b = genProb()
			if !b {
				fmt.Println("Game Over Brewski!\n-Computer")
				i += 2
			}
		}
	} else {
		fmt.Println("Lame dude")
		fmt.Println("-Computer")
	}
}
