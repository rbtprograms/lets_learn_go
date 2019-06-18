package main

import "fmt"

func main() {
	fmt.Println("hi")
	test()
	test2("HELLO")
	s1 := woo("Bobby")
	fmt.Println(s1)
	x, y := mouse("Sophia", "Fujiki")
	fmt.Println(x)
	fmt.Println(y)
}

func test() {
	fmt.Println("TEST")
}

func test2(s string) {
	fmt.Println("ARG: ", s)
}

func woo(s string) string {
	return fmt.Sprint("Hello from woo, ", s)
}

func mouse(first string, last string) (string, bool) {
	a := fmt.Sprint(first, last, `, says "Hello"`)
	b := false
	return a, b
}
