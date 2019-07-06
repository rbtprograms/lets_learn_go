package main

import (
	"fmt"
)

//channels are REFERENCE type, they reference the same data
func main() {
	c := make(chan int)

	//can pass down a bi-directional channel and convert it into
	//mono directional in the functions and it works
	go sendToChan(c)

	//this one is not a go routine because it will block the thread until
	//it does its thing, could use wait groups but this is a little easier
	// go receiveFromChan(c)
	receiveFromChan(c)

	fmt.Println("EXITING")
}

//send only channel
func sendToChan(c chan<- int) {
	c <- 42
}

//receive only channel
func receiveFromChan(c <-chan int) {
	fmt.Println(<-c)
}