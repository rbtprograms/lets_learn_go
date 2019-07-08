package main

import "fmt"

func main() {
	exercise1()
	exercise2()
	exercise3()
	// exercise4()
	// exercise5()
	// exercise6()
	// exercise7()
}

func exercise1() {
	fmt.Println("EXERCISE 1:")
	c := make(chan int)
	c2 := make(chan int, 1)

	go func() {
		c <- 42
	}()
	c2 <- 43

	fmt.Println(<-c)
	fmt.Println(<-c2)
}

func exercise2() {
	fmt.Println("EXERCISE 2:")

	cs := make(chan int)

	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)

	cr := make(chan int)

	go func() {
		cr <- 42
	}()
	fmt.Println(<-cr)

	fmt.Printf("------\n")
	fmt.Printf("cr\t%T\n", cr)
}

func gen(c chan<- int) {

	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
}

func receive(c chan int) {
	for v := range c {
		fmt.Println(v)
	}
}

func exercise3() {
	c := make(chan int)

	fmt.Println("EXERCISE 3:")
	go gen(c)
	receive(c)

	fmt.Println("about to exit")
}

func exercise4() {
	s := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	s = append(s, 52)
	fmt.Println(s)
	s = append(s, 53, 54, 55)
	fmt.Println(s)
	s = append(s, []int{56, 57, 58, 59}...)
	fmt.Println(s)
}

func exercise5() {
	s := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	s = append(s[:3], s[6:]...)
	fmt.Println(s)
}

func exercise6() {
	s := make([]string, 50, 50)
	s = []string{}
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))
}
func exercise7() {
	s := [][]string{[]string{"Bobby", "Sophia", "yith"}, []string{"Cait", "Mom", "Dad"}}
	fmt.Println(s)

	for i, v := range s {
		fmt.Println("item: ", i)
		for j, v := range v {
			fmt.Printf("\tindex: %v\tvalue: %v", j, v)
		}
	}
}
