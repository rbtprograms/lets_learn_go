package main

import "fmt"

func main() {
	exercise1()
	exercise2()
}

func exercise1() {
	x := 42
	fmt.Println(x)
	fmt.Println(&x)
}

type person struct{
	first string
	last string
}
func changeMe(p *person) {
	(*p).first = "CHANGED"
}

func exercise2() {
	p := person{
		first: "Bobby",
		last: "Bubbero",
	}
	fmt.Println(p)
	changeMe(&p)
	fmt.Println(p)
}