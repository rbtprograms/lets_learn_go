package main

import (
	"fmt"
	"runtime"
	"sync"
)


//run this code with the lines commented in
//run this with go run -race main.go to see some debugging info
func main() {
	fmt.Println("CPUS:", runtime.NumCPU())
	fmt.Println("FIRST Goroutines:", runtime.NumGoroutine())

	counter := 0

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			//this will yield the thread, honestly i need to read more about it
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
			// fmt.Println("SECOND Goroutines:", runtime.NumGoroutine())
		}()
		// fmt.Println("SECOND Goroutines:", runtime.NumGoroutine())
	}
	fmt.Println("SECOND Goroutines:", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("COUNT:", counter)
}
