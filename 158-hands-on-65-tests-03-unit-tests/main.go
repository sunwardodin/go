package main

import (
	"fmt"
	"math"
)

func main() {

	x := doSquare(2, 3, square)

	fmt.Println(x)

}

func doSquare(a float64, b float64, f func(float64, float64) float64) float64 {
	return f(a, b)
}

func square(a float64, b float64) float64 {
	total := math.Pow(a, 2) + math.Pow(b, 2)
	return total
}
