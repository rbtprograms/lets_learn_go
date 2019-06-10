package main

import "fmt"

//all variables, experissions, and statements in go
//that are declared without value are auto set
//to that types 0 value: 0, false, nil, '', etc.
var z int

func main() {
	//these declarations are the same
	var x = 42
	fmt.Println(x)
	y := 90
	fmt.Println(y)
	fmt.Println(z)

	scope()
}

func scope() {
	fmt.Println(z)
}