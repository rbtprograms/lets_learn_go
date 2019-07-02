package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

type been struct {
	person
	been bool
}

func main() {
	b1 := been{
		person: person{
			first: "Bobby",
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
	fmt.Println(b1.first)
	fmt.Println(b1.person.last)
	fmt.Println(b2)
	fmt.Println(b2.last)
}
