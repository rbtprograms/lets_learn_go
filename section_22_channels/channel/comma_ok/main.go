package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go func(){
		c <- 100
		close(c)
	}()
	//first will ahve 100 true, second will have 0 false
	//can use the bool off the channel to do stuff
	v, ok := <-c
	fmt.Println("off the chan:", v, ok)
	v, ok = <-c
	fmt.Println("off the chan:", v, ok)
}

