package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

//be careful with package atomic, it is very low level
//and takes a lot of care to use correctly
func main() {
	fmt.Println("CPUS:", runtime.NumCPU())
	fmt.Println("FIRST Goroutines:", runtime.NumGoroutine())

	var counter int64

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Println("Counter\t", atomic.LoadInt64(&counter))
			runtime.Gosched()
			wg.Done()
			// fmt.Println("SECOND Goroutines:", runtime.NumGoroutine())
		}()
		// fmt.Println("SECOND Goroutines:", runtime.NumGoroutine())
	}
	fmt.Println("SECOND Goroutines:", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("COUNT:", counter)
}
