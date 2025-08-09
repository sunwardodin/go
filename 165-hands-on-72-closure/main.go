package main

import (
	"fmt"
	"math"
)

func close(a float64) func() float64 {
	var x float64
	return func() float64 {
		x++
		return math.Pow(a, x)
	}
}

func main() {
	x := close(2)
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
}
