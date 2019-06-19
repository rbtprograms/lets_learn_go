package main

import "fmt"

//theres a lot going on here,
//values can be multiple types
//interfaces allow us to mesh functionalities of different types

type person struct {
	first string
	last  string
	age   int
}

type been struct {
	person
	been bool
}

func (s been) speak() {
	fmt.Println("I am", s.first, s.last, "and I am a been")
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last, "The human")
}

type human interface {
	speak()
}

func test(h human) {
	switch h.(type) {
	case person:
		fmt.Println("Test function was called with the follwing: ", h.(person))
	case been:
		fmt.Println("Test function was called with the follwing: ", h.(been).first)
	}
}

func main() {
	b1 := been{
		person: person{
			first: "Bubby",
			last:  "Bubbero",
			age:   29,
		},
		been: true,
	}
	b2 := been{
		person: person{
			first: "Subby",
			last:  "Subbero",
			age:   27,
		},
		been: true,
	}
	p1 := person{
		first: "Dr.",
		last: "Yes",
		age: 50,
	}
	//embedded fields are promoted up but can still be 
	//accessed via their type
	fmt.Println(b1)
	fmt.Println(b1.first)
	fmt.Println(b1.person.last)
	fmt.Println(b2)
	fmt.Println(b2.last)
	fmt.Println(p1)

	//these can accept all these despite different
	//typing cause they are all type human
	test(b1)
	test(b2)
	test(p1)
}
