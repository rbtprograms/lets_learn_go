package main

import (
	"fmt" 
	"math"
)

func main() {
	exercise1()
	exercise2()
	exercise3()
	exercise4()
	exercise5()
	exercise6()
	exercise7()
	exercise8()
	exercise9()
	exercise10()
}

func exercise1() {
	foo := func() int {
		return 5
	}
	bar := func() (int, string) {
		return 55, "hello"
	}
	r1 := foo()
	r2a, r2b := bar()
	fmt.Println("r1", r1)
	fmt.Println("r2a", r2a)
	fmt.Println("r2b", r2b)
}
func exercise2() {
	foo := func(x ...int) int {
		var sum int
		for _, v := range x {
			sum += v
		}
		return sum
	}
	bar := func(x []int) int {
		var sum int
		for i := range x {
			sum += x[i]
		}
		return sum
	}
	vals := []int{1, 2, 3, 4, 5}
	fmt.Println("foo:", foo(vals...))
	fmt.Println("bar:", bar(vals))
}

func exercise3() {
	defer fmt.Println("FIRST PRINTLN")
	fmt.Println("SECOND PRINTLN")
}

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Printf("\nHi my name is %v %v. I am %v years old.\n", p.first, p.last, p.age)
}

func exercise4() {
	p := person{
		first: "Bobby",
		last:  "Thompson",
		age:   29,
	}
	p.speak()
}

type square struct {
	length float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return math.Pow(s.length, 2)
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(2, 2)
}

type shape interface {
	area() float64
}

func info(s shape){
	fmt.Println(s.area())
}

func exercise5() {
	s := square{
		length: 5,
	}
	c := circle{
		radius: 6.5,
	}
	info(s)
	info(c)
	
}

var temp = 5

func exercise6() {
	func() {
		fmt.Println("an unnamed function is called")
	}()
}

var g = func() {
	fmt.Println("function called from outside the amin block")
}
func exercise7() {
	f := func() {
		fmt.Println("Variable expression function was called")
	}
	f()
	g()
}
func returnsFunc() func() string {
	return func() string {
		return "hi"
	}
}
func exercise8() {
	h := returnsFunc()
	res := h()
	fmt.Println(res)
}

func foo(f func(xi []int) int, ii []int) int {
	n := f(ii)
	n++
	return n
}
func exercise9() {
	g := func(xi []int) int {
		if len(xi) == 0 {
			return 0
		}
		if len(xi) == 1 {
			return xi[0]
		}
		return xi[0] + xi[len(xi) - 1]
	}
	x := foo(g, []int{1,2,3,4,5,6})
	fmt.Println(x)
}

func exercise10() {
	increment := func() func() int {
		x := 0
		return func() int {
			x++
			return x
		}
	}
	a := increment()
	b := increment()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
}