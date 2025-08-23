package main

import (
	"fmt"
	"math/rand"
)

// Math Game

func add() {
	score := 0
	min := 10
	max := 100

	for i := 0; i <= 1; {
		fmt.Println("----------------")
		for j := 0; j <= 1; j++ {
			n1 := min + rand.Intn(max+min)
			n2 := min + rand.Intn(max+min)
			answer := 0
			cAnswer := n1 + n2

			fmt.Printf("%v + %v = ", n1, n2)
			fmt.Scanln(&answer)
			if answer != cAnswer {
				i += 2
				break
			}
			score++
		}
		min += 250
		max += 250
	}

	fmt.Println("\nYour score was:", score)
	fmt.Println("\nGame Over Brewski!\n-Computer")
}

func subtract() {
	score := 0
	min := 10
	max := 100

	for i := 0; i <= 1; {
		fmt.Println("----------------")
		for j := 0; j <= 1; j++ {
			n1 := min + rand.Intn(max+min)
			n2 := min + rand.Intn(max+min)
			answer := 0
			cAnswer := n1 - n2

			fmt.Printf("%v - %v = ", n1, n2)
			fmt.Scanln(&answer)
			if answer != cAnswer {
				i += 2
				break
			}
			score++
		}
		min += 250
		max += 250
	}

	fmt.Println("\nYour score was:", score)
	fmt.Println("\nGame Over Brewski!\n-Computer")
}

func multiply() {
	score := 0
	min := 10
	max := 100

	for i := 0; i <= 1; {
		fmt.Println("----------------")
		for j := 0; j <= 1; j++ {
			n1 := min + rand.Intn(max+min)
			n2 := min + rand.Intn(max+min)
			answer := 0
			cAnswer := n1 * n2

			fmt.Printf("%v * %v = ", n1, n2)
			fmt.Scanln(&answer)
			if answer != cAnswer {
				i += 2
				break
			}
			score++
		}
		min += 250
		max += 250
	}

	fmt.Println("\nYour score was:", score)
	fmt.Println("\nGame Over Brewski!\n-Computer")
}

func divide() {
	score := 0

	for i := 0; i < 1; {
		d1 := 50 + rand.Intn(100-50)
		d2 := 10 + rand.Intn(50-10)
		var dAnswer int
		cDAnswer := d1 / d2

		fmt.Printf("What is %v / %v\n", d1, d2)
		fmt.Scanln(&dAnswer)
		if dAnswer != cDAnswer {
			i++
			break
		}
		score++
	}
	fmt.Println("\nYour score was:", score)
	fmt.Println("\nGame Over Brewski!\n-Computer")
}

func main() {
	var a string
	fmt.Print("Ready for a problem? ")
	fmt.Scanln(&a)

	if a == "yes" || a == "Yes" {
		for i := 0; i < 1; {

			var d string
			fmt.Print("What mode? ")
			fmt.Scanln(&d)

			switch d {
			case "addition":
				add()
			case "subtraction":
				subtract()
			case "multiplication":
				multiply()
			case "division":
				divide()
			default:
				fmt.Println("That isn't a mode. Try again.")
			}
			var b string
			fmt.Print("Want to play again? ")
			fmt.Scanln(&b)
			if b == "no" {
				i += 2
			}
		}
	} else {
		fmt.Println("Lame dude")
		fmt.Println("-Computer")
	}
}

/*
var hE string
	fmt.Printf("hard or easy? ")
	fmt.Scanln(&hE)

	if hE == "hard" {
		for i := 0; i < 1; {
			n1 := rand.Intn(10000)
			n2 := rand.Intn(10000)
			answer := 0
			cAnswer := n1 + n2

			fmt.Printf("What is %v + %v\n", n1, n2)
			fmt.Scanln(&answer)
			if answer != cAnswer {
				i++
				break
			}
			score++
		}
	} else if hE == "easy" {
		for i := 0; i < 1; {
			n1 := rand.Intn(100)
			n2 := rand.Intn(100)
			answer := 0
			cAnswer := n1 + n2

			fmt.Printf("What is %v + %v\n", n1, n2)
			fmt.Scanln(&answer)
			if answer != cAnswer {
				i++
				break
			}
			score++
		}
	}
*/
