package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	testFunction()

	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

func testFunction() {
	fmt.Println("Inside of testFunction")
}
