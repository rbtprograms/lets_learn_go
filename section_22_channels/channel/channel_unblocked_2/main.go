package main

import (
	"fmt"
	"runtime"
)

//this one runs because it is making a buffered channel. I am specifiying
//when i make the channel that ONE value is allowed to sit in there, so nothing
//has to block to put a value in there
func main() {
	fmt.Println("GO ROUTINES", runtime.NumGoroutine())
	c := make(chan int, 1)

	c <- 42
	//comment in below line and try to run
	// c <- 43
	//The code will be blocked until you take a value off the channel

	fmt.Println(<-c)
}
