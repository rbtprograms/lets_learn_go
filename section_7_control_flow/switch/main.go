package main

import "fmt"

//no while loops in go
func main() {
	switch "test" {
	case "false":
		fmt.Println("false")
	case "true":
		fmt.Println("true")
	case "Howdy", "Hi", "test":
		fmt.Println("Hi there")
	default:
		fmt.Println("true")
	}
}
