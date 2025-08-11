package main

import "fmt"

func intDelta(n *int) {
	*n = 43
}

func sliceDelta(ii []int) {
	ii[0] = 99
}

func mapDelta(md map[string]int, s string) {
	md[s] = 33
}

func main() {
	x := 42
	fmt.Println(x)
	intDelta(&x)
	fmt.Println(x)

	xi := []int{1, 2, 3, 4}
	fmt.Println(xi)
	sliceDelta(xi) // Since a slice is already a reference type, when this function is called it will change the slice directly.
	fmt.Println(xi)

	m := make(map[string]int)
	m["James"] = 32
	fmt.Println(m["James"])
	mapDelta(m, "James") // This is the exact same idea as the slice
	fmt.Println(m["James"])
}
