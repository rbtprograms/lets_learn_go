package main

import "fmt"

//no while loops in go
func main() {
	for i := 0; i <= 10; i++ {
		for j := 0; j <= 10; j++ {
			fmt.Printf("Outer loop: %v\t and the Inner: %v\n", i, j)
		}
	}

}
