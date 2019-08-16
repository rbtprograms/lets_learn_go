package main

import "fmt"

//returns hello world
func Hello(name string) string {
	return "hello " + name
}

func main() {
	fmt.Println(Hello("Bobby"))
}