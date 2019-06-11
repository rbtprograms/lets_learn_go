package main

import "fmt"

var y string

func main() {
	fmt.Println("value of y: ", y)
	fmt.Printf("%T\n", y)
	y = "Hi there"
}