package main

import (
	"fmt"
	"encoding/json"
)

type person struct{
	First string
	Last string
	Age int
}

func main() {
	p1 := person{
		First: "Bobby",
		Last: "Thompson",
		Age: 29,
	}
	p2 := person{
		First: "Sophia",
		Last: "Fujiki",
		Age: 27,
	}
	people := []person{p1, p2,}
	fmt.Println(people)
	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println("ERROR:", err)
	}
	fmt.Println(string(bs))
}
