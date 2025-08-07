package main

import "fmt"

func main() {
	f := incrementor(5)
	fmt.Println(f()) // 1
	fmt.Println(f()) // 2
	fmt.Println(f()) // 3
	fmt.Println(f()) // 4
	fmt.Println(f()) // 5
	fmt.Println(f()) // 6

	g := incrementor(25) // Since this is a different variable, it is a different memory location altogether. Meaning, it will print out its own set of numbers
	fmt.Println(g())     // 1
	fmt.Println(g())     // 2
	fmt.Println(g())     // 3
	fmt.Println(g())     // 4
	fmt.Println(g())     // 5
	fmt.Println(g())     // 6
}

func incrementor(a int) func() int {
	return func() int {
		a++
		return a
	}
}
