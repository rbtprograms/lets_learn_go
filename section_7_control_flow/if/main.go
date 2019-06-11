package main

import "fmt"

//no while loops in go
func main() {

	//init statement limits the scope of x
	if x := 42; x == 3 {
		fmt.Println("true")
	}
	if x == false {
		fmt.Println("false")
	}
	if 2 == 2 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
