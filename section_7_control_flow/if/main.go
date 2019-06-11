package main

import "fmt"

//no while loops in go
func main() {

	//init statement limits the scope of x
	if x := 42; x == 3 {
		fmt.Println("true")
	}

	y := 42
	if y == 40 {
		fmt.Println("40")
	} else if y == 41 {
		fmt.Println("41")
	} else {
		fmt.Println("neither 40 nor 41")
	}
}
