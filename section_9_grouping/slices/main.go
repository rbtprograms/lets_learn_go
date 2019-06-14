package main

import (
	"fmt"
)

//arrays are not used much in go, use slices instead
//slices have dynamic sizes
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

	//using make
	//last item is max amount of elements
	x = make([]int, 10, 12)
	x[0] = 1
	x[9] = 99
	//can keep appending until cap is hit
	//go will auto extend cap if you go over
	x = append(x, 333)
	x = append(x, 333)
	x = append(x, 333)
	fmt.Println(x)
	fmt.Println(len(x))
	//cap shows how many elements it can hold
	fmt.Println(cap(x))

	bt := []string{"Bobby", "T", "Chocolate", "Honey"}
	fmt.Println(bt)
	sf := []string{"Subby", "Fujiki", "Vanilla", "Strawberry"}
	fmt.Println(sf)

	xp := [][]string{bt, sf}
	fmt.Println(xp)

	//This is kind of like how objects in JS point to the same object,
//so copying them is tricky. Slices will point to the same underlying
//array, so if one of your slices changes the underlying array it will
//affect another slice using data from that same array. thats
//why b is changed when c is changed
b := []int{65,43,6,88,32,56,9,32,4,67}
fmt.Println(b)
c := append(b[0:2], b[5:]...) //new underlying array stores the value of new slice

fmt.Println(b)
fmt.Println(c)
}

//a slice allows you to group together values
//of the same type
