package main

import (
	"fmt"
	"sync"
)

func main() {
	even := make(chan int)
	odd := make(chan int)
	fanIn := make(chan int)

	go send(even, odd)
	go receive(even, odd, fanIn)

	for v := range fanIn {
		fmt.Println(v)
	}

	fmt.Println("exiting...")
}

func receive(e, o <-chan int, fi chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func(){
		for v:= range e {
			fi <- v
		}
		wg.Done()
	}()

	go func() {
		for v:= range o {
			fi <- v
		}
		wg.Done()
	}()

	wg.Wait()
	close(fi)
}

func send(e, o chan<- int) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(e)
	close(o)
}
