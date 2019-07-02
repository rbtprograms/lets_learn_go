package main

import "fmt"

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
	fmt.Println("I am", s.first, s.last)
}

func main() {
	b1 := been{
		person: person{
			first: "Bouby",
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
	//embedded fields are promoted up but can still be
	//accessed via their type
	fmt.Println(b1)
	fmt.Println(b2)
	b1.speak()
	b2.speak()
}
