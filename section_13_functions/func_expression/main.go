package main

import (
	"fmt"
)

//functions are first class citizens;
//can do anything other types can do
func main() {
	f := func(){
		fmt.Println("first func expression")
	}
	f()

	f2 := func(x int){
		fmt.Println("PARAM:", x)
	}
	f2(1984)
}