package main

import (
	"fmt"
)

func main() {
	test()

	func(x int){
		fmt.Println("anonymous function lets go")
		fmt.Println("PASSED PARAM:", x)
	}(42)
}

func test() {
	fmt.Println("hi")
}
