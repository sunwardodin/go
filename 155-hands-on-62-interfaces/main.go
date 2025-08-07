package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
	width  float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.length * s.width
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

type shape interface {
	area() float64
}

func info(s shape) float64 {
	return s.area()
}

func main() {
	s := square{
		length: 55.5,
		width:  66.6,
	}
	c := circle{
		radius: 45.5,
	}

	fmt.Println(info(s))
	fmt.Println(info(c))
}
