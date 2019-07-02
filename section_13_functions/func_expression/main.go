package main

import (
	"fmt"
)

//functions are first class citizens:
//can do anything other types can do
func main() {
	f := func() {
		fmt.Println("first func expression")
	}
	f()

	f2 := func(x int) {
		fmt.Println("PARAM:", x)
	}
	f2(1984)
	fmt.Println(foo())

	fmt.Println(bar()())

	x := bar()
	fmt.Printf("\n%T\n", x)
	y := x()
	fmt.Printf("\n%T\n%v\n", y, y)
}

func foo() string {
	s := "Hello world"
	return s
}

func bar() func() int {
	return func() int {
		return 451
	}
}
