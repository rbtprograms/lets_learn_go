package main

import "fmt"

var y = 42

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Printf("%b\n", y)
	fmt.Printf("%x\n", y)
	fmt.Printf("%#x\n", y)
	y = 99
	fmt.Println(y)
	fmt.Printf("%T\t%b\t%x\n", y, y, y)

	//creates a string
	s := fmt.Sprintf("%T\t%b\t%x\t", y, y, y)
	fmt.Println(s)

	fmt.Printf("%v", y)
}