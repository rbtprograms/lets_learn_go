package main

import "fmt"

func main() {
	//runs second, at end of containing function
	defer test()
	test2()
}

func test() {
	fmt.Println("TEST")
}

func test2() {
	fmt.Println("SECOND TEST")
}
