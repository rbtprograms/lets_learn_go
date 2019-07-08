package main

import (
	"fmt"
	"runtime"
)

//channels MUST receive and pass data at the same time,
//channels will block execution
//this code doesnt run because the channel is blocking
func main() {
	fmt.Println("GO ROUTINES", runtime.NumGoroutine())
	c := make(chan int)

	c <- 42

	fmt.Println(<-c)
}
