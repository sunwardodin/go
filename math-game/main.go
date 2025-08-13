package main

import (
	"fmt"
	"math/rand"
)

// Math Game
func genProb() bool {
	sl := []string{"+", "-", "/", "*", "%"}
	randomOp := sl[rand.Intn(4)]

	n1 := rand.Intn(100)
	n2 := rand.Intn(100)
	answer := 0
	cAnswer := 0

	d1 := 50 + rand.Intn(100-50)
	d2 := 10 + rand.Intn(50-10)
	var dAnswer int
	var cDAnswer int

	switch randomOp {
	case "+":
		cAnswer = n1 + n2
		fmt.Printf("What is %v + %v\n", n1, n2)
		fmt.Scanln(&answer)
		if answer != cAnswer {
			return false
		}
	case "-":
		cAnswer = n1 - n2
		fmt.Printf("What is %v - %v\n", n1, n2)
		fmt.Scanln(&answer)
		if answer != cAnswer {
			return false
		}
	case "*":
		cAnswer = n1 * n2
		fmt.Printf("What is %v * %v\n", n1, n2)
		fmt.Scanln(&answer)
		if answer != cAnswer {
			return false
		}

	case "/":
		cDAnswer = d1 / d2
		fmt.Printf("What is %v / %v\n", d1, d2)
		fmt.Scanln(&dAnswer)
		if dAnswer != cDAnswer {
			return false
		}
	case "%":
		cDAnswer = d1 % d2
		fmt.Printf("What is %v '%' %v \n", d1, d2)
		fmt.Scanln(&dAnswer)
		if dAnswer != cDAnswer {
			return false
		}

	}
	return true
}

func hardCore() bool {
	n1 := rand.Intn(1000)
	n2 := rand.Intn(1000)
	answer := 0
	cAnswer := 0

	cAnswer = n1 * n2
	fmt.Printf("What is %v * %v\n", n1, n2)
	fmt.Scanln(&answer)

	return answer == cAnswer
}

func main() {
	var a string
	var b bool
	i := 0

	fmt.Print("Ready for a problem?: ")
	fmt.Scanln(&a)
	switch a {
	case "yes":
		score := 0
		for i < 1 {
			b = genProb()
			if !b {
				fmt.Println("\nYour score was:", score)
				fmt.Println("\nGame Over Brewski!\n-Computer")
				i += 1
			}
			score += 1
		}
	case "hardcore":
		score := 0
		for i < 1 {
			b = hardCore()
			if !b {
				fmt.Println("\nYour score was:", score)
				fmt.Println("\nGame Over Brewski!\n-Computer")
				i += 1
			}
			score += 1
		}
	default:
		fmt.Println("Lame dude")
		fmt.Println("-Computer")
	}
}
