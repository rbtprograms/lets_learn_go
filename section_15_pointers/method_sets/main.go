package main

import (
	"fmt"
	"math"
)

//when defining methods, non pointer receivers can deal
//with pointers or non pointers. Pointer receivers
//can only be utilized with pointers, not non-pointer values

type square struct {
	length float64
}

type circle struct {
	radius float64
}

func (s *square) area() float64 {
	return math.Pow(s.length, 2)
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(2, 2)
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	s := square{
		length: 5,
	}
	c := circle{
		radius: 6.5,
	}
	//can be pointers for non-pointer receivers
	info(&s)
	info(&c)

}
