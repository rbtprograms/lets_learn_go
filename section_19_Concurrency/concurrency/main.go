package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUS\t\t", runtime.NumCPU())
	fmt.Println("GOROUTINES\t", runtime.NumGoroutine())

	//wait for 1 event
	wg.Add(1)
	//start concurrent routine
	go foo()
	fmt.Println("GOROUTINES\t", runtime.NumGoroutine())
	bar()
	//wait before exiting main function
	wg.Wait()
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}
	//this will tick off 1 from the wait
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}
