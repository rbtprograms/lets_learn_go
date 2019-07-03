package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 20; i++ {
			c <- i
		}
		//c must be closed, it will deadlock the program as it is waiting for more values
		close(c)
	}()

	//ranging through the channel pulls off all values and closes the channel
	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("ENDING PROGRAM")
}
