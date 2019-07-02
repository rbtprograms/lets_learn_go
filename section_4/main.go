package main

import "fmt"

func main() {
	exercise1()
	exercise2()
	exercise3()
	exercise4()
	exercise5()
}

func exercise1() {
	x := 42
	y := "James Bond"
	z := true
	fmt.Println(x, y, z)
}

func exercise2() {
	var x int
	var y string
	var z bool
	fmt.Printf("%v\n", x)
	fmt.Printf("%v\n", y)
	fmt.Printf("%v\n", z)
}

func exercise3() {
	x := 42
	y := "JamesBond"
	z := true
	s := fmt.Sprintf("%v\t%v\t%v\t", x, y, z)
	fmt.Println(s)
}

type newNum int

var x newNum

func exercise4() {
	fmt.Printf("%T\t", x)
	x = 42
	fmt.Println(x)
}

var y int

func exercise5() {
	y = 100
	y = int(x)
	fmt.Printf("%T\t", y)
	fmt.Printf("%v\t", y)
}
