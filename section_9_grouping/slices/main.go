package main

import (
	"fmt"
)

//arrays ar enot used much in go, use slices instead
func main() {
	//making a slice
	//x := type{values} //composite literal
	x := []int{4, 5, 6, 7, 8}
	fmt.Println("value: ", x)
	fmt.Println("length: ", len(x))

	//iterate over a slice
	for i, v := range x {
		fmt.Println(i, v)
	}

	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}

	//slicing a slice
	fmt.Println(x[1:])
	fmt.Println(x[2:4])

	//append to a slice
	y := []int{9,10,11}
	x = append(x, y...)
	//						^^^^go's spread operator
	fmt.Println(x)

	//delete from a slice
	x = append(x[0:2], x[4:]...)
	fmt.Println(x)
}

//a slice allows you to group together values
//of the same type
