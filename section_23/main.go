package main

import "fmt"

func main() {
	exercise1()
	exercise2()
	exercise3()
	exercise4()
	exercise5()
	exercise6()
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

func exercise3() {
	c := make(chan int)

	fmt.Println("EXERCISE 3:")
	go func(c chan<- int) {

		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}(c)
	func(c chan int) {
		for v := range c {
			fmt.Println(v)
		}
	}(c)

	fmt.Println("about to exit")
}

func exercise4() {
	fmt.Println("EXERCISE 4:")

	q := make(chan int)
	c := func(q chan<- int) <-chan int {
		c := make(chan int)

		go func() {

			for i := 20; i < 30; i++ {
				c <- i
			}
			q <- 1
			close(c)
		}()
		return c
	}(q)

	func(c <-chan int, q chan int) {
		for {
			select {
			case v := <-c:
				fmt.Println("value pulled off channel:", v)
			case <-q:
				return
			}
		}
	}(c, q)

	fmt.Println("about to exit")
}

func exercise5() {
	fmt.Println("EXERCISE 5:")

	c := make(chan int)

	go func(c chan<- int) {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}(c)

	func(c <-chan int) {
		for v := range c {
			fmt.Println(v)
		}
	}(c)
}

func exercise6() {
	c := make(chan int)
	for i := 0; i < 10; i++ {
		go func(c chan<- int) {
			for j := 0; j < 10; j++ {
				c <- j
			}
		}(c)
	}
	for i := 0; i < 100; i++ {
		fmt.Println(i, <-c)
	}

}
