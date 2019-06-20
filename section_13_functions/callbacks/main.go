package main

import (
	"fmt"
)

//functions are first class citizens:
//can do anything other types can do
func main() {
	ii := []int{1,2,3,4,5,6,7,8,9}
	s := sum(ii...)
	eS := evenSum(sum, ii...)
	oS := oddSum(sum, ii...)
	fmt.Println("all numbers sum", s)
	fmt.Println("even Number sum", eS)
	fmt.Println("even Number sum", oS)
}

func sum(xi ...int) int {
	fmt.Printf("%T\n%v\n", xi, xi)
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func evenSum(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v % 2 == 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}

func oddSum(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v % 2 != 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}