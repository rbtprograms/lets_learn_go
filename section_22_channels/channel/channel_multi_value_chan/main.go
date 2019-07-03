package main

import (
	"fmt"
	"runtime"
)

//buffered channels are generally considered bad practice
func main() {
	fmt.Println("GO ROUTINES", runtime.NumGoroutine())
	c := make(chan int, 2)

	c <- 42
	c <- 43

	//we are taking values off the channel as we log them.
	//this prints 42
	fmt.Println(<-c)
	//this prints 43
	fmt.Println(<-c)
}
