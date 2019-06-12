package main

import (
	"fmt"
)

//arrays ar enot used much in go, use slices instead
func main() {
	var x [5]int
	x[3] = 42
	fmt.Println(x)
}
