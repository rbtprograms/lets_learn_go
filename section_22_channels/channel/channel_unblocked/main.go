package main

import (
	"fmt"
	"runtime"
)

//this version runs because the go routine goes off, puts the data
//in the channel and gets blocked, but brings it back to the main
//function and they link up and it works!
func main() {
	fmt.Println("GO ROUTINES", runtime.NumGoroutine())
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}
