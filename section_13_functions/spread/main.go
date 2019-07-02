package main

import "fmt"

func main() {
	test()
	xi := []int{2, 3, 4, 5, 6, 7, 8, 9}
	res := test2(xi...)
	fmt.Println(res)
}

func test() {
	fmt.Println("TEST")
}

func test2(x ...int) int {
	//^^^this makes a slice
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	sum := 0
	for _, v := range x {
		sum += v
		fmt.Println("item: ", v, " and sum: ", sum)
	}
	return sum
}
