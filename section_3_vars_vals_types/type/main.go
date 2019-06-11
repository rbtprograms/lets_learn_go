package main

import "fmt"

var y = 8
var z = "hi"
var a = `honk if you like dying
and being dead`

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", z)
	fmt.Println(a)
}