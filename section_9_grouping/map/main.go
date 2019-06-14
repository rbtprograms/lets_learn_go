package main

import (
	"fmt"
)

//This is kind of like how objects in JS point to the same object,
//so copying them is tricky. Slices will point to the same underlying
//array, so if one of your slices changes the underlying array it will
//affect another slice using data from that same array.
func main() {
	x := []int{65,43,6,88,32,56,9,32,4,67}
	fmt.Println(x)
	y := append(x[0:2], x[5:]...) //new underlying array stores the value of new slice

	fmt.Println(x)
	fmt.Println(y)
}
