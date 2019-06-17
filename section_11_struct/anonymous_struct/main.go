package main

import (
	"fmt"
)

func main() {
	b1 := struct {
		first string
		last  string
		age   int
	}{
		first: "Bobby",
		last:  "Bubbero",
		age:   29,
	}
	fmt.Println(b1)
	fmt.Println(b1.first)
}
