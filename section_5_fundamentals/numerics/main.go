package main

import (
	"fmt"
	"runtime"
)

var x int
var y float64
var z int8 = -128

func main() {
	x = 100
	y = 100.91827
	fmt.Printf("%v\t%T\n", x, x)
	fmt.Printf("%v\t%T\n", y, y)

	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}
