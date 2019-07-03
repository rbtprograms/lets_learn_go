package main

import (
	"fmt"
	"sync"
)

func main() {
	exercise1()
	exercise2()
}

func exercise1() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("FIRST GO ROUTINE")
		wg.Done()
	}()
	go func() {
		fmt.Println("SECOND GO ROUTINE")
		wg.Done()
	}()
	wg.Wait()
}

type person struct {
	first string
	last  string
}

type dog struct {
	name  string
	color string
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Printf("My name is %v %v\n", p.first, p.last)
}

func (d *dog) speak() {
	fmt.Println("WOOF WOOF I AM A DOG")
}

func saySomething(h human) {
	h.speak()
}

func exercise2() {
	p := person{
		first: "Bubby",
		last:  "Bubbero",
	}
	d := dog{
		name:  "Fido",
		color: "brown",
	}
	saySomething(&p)
	saySomething(&d)
}
