package main

import "fmt"

func main() {
	test()
}

func test() {
	fmt.Println("TEST")
	res := test2(2,3,4,5,6,7,8,9)
	fmt.Println(res)
}

func test2(x ...int) int {
					 //^^^this makes a slice
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for _, v := range x {
		sum += v
		fmt.Println("item: ", v, " and sum: ", sum)
	}
	return sum
}