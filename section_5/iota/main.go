package main

import "fmt"

//iota are predetermined type, they start at 0
//and auto increment but only in the same const block

const (
	a = iota //0
	b = iota //1
	c = iota //2
)

const (
	d = iota 	//0
	e					//1
	f					//2
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
}
