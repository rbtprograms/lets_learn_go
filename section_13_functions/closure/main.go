package main

import (
	"fmt"
)

func main() {
	a := increment()
	b := increment()
	fmt.Println("a", a())
	fmt.Println("a", a())
	fmt.Println("b", b())
	fmt.Println("b", b())
	fmt.Println("b", b())
	fmt.Println("b", b())
}

func increment() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
