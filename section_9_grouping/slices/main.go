package main

import (
	"fmt"
)

//arrays ar enot used much in go, use slices instead
func main() {
	//x := type{values} //composite literal
	x := []int{4, 5, 6, 7, 8}
	fmt.Println("value: ", x)
	fmt.Println("length: ", len(x))

	for i, v := range x {
		fmt.Println(i, v)
	}
}

//a slice allows you to group together values
//of the same type
