package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUS:", runtime.NumCPU())
	fmt.Println("FIRST Goroutines:", runtime.NumGoroutine())

	counter := 0

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	//mutex locks down variables so you cannot have multiple
	//go routines using
	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			//locks down variables
			mu.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			//opens them up
			mu.Unlock()
			wg.Done()
		}()
	}
	fmt.Println("SECOND Goroutines:", runtime.NumGoroutine())
	wg.Wait()
	//run this and compare the count against the race_condition code
	fmt.Println("COUNT:", counter)
}
