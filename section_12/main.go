package main

import "fmt"

func main() {
	exercise1()
	exercise2()
	exercise3()
	// exercise4()
}

type person struct{
	first string
	last string
	faveFlavors []string
}

func exercise1() {
	p1 := person{
		first: "Bobby",
		last: "Bubbero",
		faveFlavors: []string{"Vanilla", "Rocky Road"},
	}
	p2 := person{
		first: "Sophia",
		last: "Subbero",
		faveFlavors: []string{"Chocolate"},
	}
	for i, v := range p1.faveFlavors {
		fmt.Println(i, v)
	}
	for i, v := range p2.faveFlavors {
		fmt.Println(i, v)
	}
}

func exercise2() {
	p1 := person{
		first: "Bobby",
		last: "Bubbero",
		faveFlavors: []string{"Vanilla", "Rocky Road"},
	}
	p2 := person{
		first: "Sophia",
		last: "Subbero",
		faveFlavors: []string{"Chocolate"},
	}

	m := map[string]person {
		p1.last:p1,
		p2.last:p2,
	}
	for _, v := range m {
		fmt.Println(v)
		fmt.Println(v.first)
		for i, val := range v.faveFlavors {
			fmt.Println(i, val)
		}
		fmt.Println("------")
	}
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