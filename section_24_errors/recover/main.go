package main

import (
	"fmt"
)

func main() {
	fmt.Print("Name:")
	a()
	b()

	y := 1
	y = c()
	fmt.Println("Y:", y)

	f()
	fmt.Println("Returned normally from f.")
}

//below will print 0 because deferred functions are
//evaluated when they are read
func a() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}

//below prints 43210 because defer functions are
//last in first out
func b() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}

//this is weird. i is the named return and
//when the function is done it will return 1 but
//the defer function will increment it and make it 2
func c() (i int) {
	defer func() { i++ }()
	return 1
}

//the two below functions are using panic
func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}
