package main

import (
	"fmt"
)

//most channels are bi-driectional: can read and write
//can specify whether or not certain channels should be one way
func main() {
	c := make(chan int)
	//can only write to channel
	receiveCh := make(chan<- int)

	go func() {
		receiveCh <- 42
	}()

	fmt.Println(<-receiveCh)
	fmt.Printf("%T\n", receiveCh)

	//can only read from channel
	sendCh := make(<-chan int)

	go func() {
		sendCh <- 42
	}()

	fmt.Println(<-sendCh)
	fmt.Printf("%T\n", sendCh)

	//specific to general doesn't work, but the inverse does
	receiveCh = c
	c = receiveCh
}
