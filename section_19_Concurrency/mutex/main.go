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
			mu.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
	}
	fmt.Println("SECOND Goroutines:", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("COUNT:", counter)
}
