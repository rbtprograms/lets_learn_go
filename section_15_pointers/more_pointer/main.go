package main

import (
	"fmt"
)

//without pointers
func main() {
	x := 0
	foo(x)
	fmt.Println(x)
}

func foo(y int) {
	fmt.Println(y)
	y = 43
	fmt.Println(y)
}

//with pointers
func main2() {
	x := 0	
	fmt.Println("first x RAW VAL", x)
	fmt.Println("first x POINTER", &x)
	foo2(&x)
	fmt.Println("second x RAW VAL", x)
	fmt.Println("second x POINTER", &x)
}

func foo2(y *int) {
	fmt.Println("FIRST Y RAW VAL", y)
	fmt.Println("FIRST Y DEREFERENCE", *y)
	*y = 43
	fmt.Println("SECOND Y RAW VAL", y)
	fmt.Println("SECOND Y DEREFERENCE", *y)
}