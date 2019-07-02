package main

import (
	"fmt"
	"sync"
)

func main() {
	exercise1()
}

func exercise1() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func(){
		fmt.Println("FIRST GO ROUTINE")
		wg.Done()
	}()
	go func(){
		fmt.Println("SECOND GO ROUTINE")
		wg.Done()
	}()
	wg.Wait()
}

type person struct {
	first string
	last  string
}

func changeMe(p *person) {
	(*p).first = "CHANGED"
}

func exercise2() {
	p := person{
		first: "Bobby",
		last:  "Bubbero",
	}
	fmt.Println(p)
	changeMe(&p)
	fmt.Println(p)
}
