package main

import "fmt"

var x bool

func main() {
	fmt.Printf("%v", x)
	x = true
	fmt.Printf("%v", x)

	a := 7
	b := 9
	if a > b {
		fmt.Println("a is greater")
	} else if b > a {
		fmt.Println("b is greater")
	} else {
		fmt.Println("They are equal")
	}
}
