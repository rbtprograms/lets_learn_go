package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	testFunction()

	n, _ := fmt.Println("Howdy", 42, true)
	fmt.Println(n)

	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
	x := 42
	fmt.Println(x);
	x = 99
	fmt.Println(x)
	var y = 101
	fmt.Println(y)
}

func testFunction() {
	fmt.Println("Inside of testFunction")
}
