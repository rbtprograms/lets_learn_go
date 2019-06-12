package main

import "fmt"

func main() {
	fmt.Println(2 == 2 && 2 == 3)
	fmt.Println(2 == 2 && 3 == 3)
	fmt.Println(2 == 4 || 3 == 3)
	fmt.Println(!false)
}
