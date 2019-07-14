package main

import (
	"fmt"
	"log"
	"encoding/json"
)

func main() {
	exercise1()
	exercise2()
	// exercise3()
	// exercise4()
	// exercise5()
	// exercise6()
}

type person struct {
	First   string
	Last    string
	Sayings []string
}

func exercise1() {
	fmt.Println("**EXERCISE 1**")
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}
	
	bs, err := json.Marshal(p1)
	if err != nil {
		log.Fatalln("Error Marshaling JSON:", err)
	}
	fmt.Println(string(bs))
}

//change to return an error
func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		return nil, fmt.Errorf("There was an error Marshaling JSON: %v", err)
	}
	return bs, nil
}

func exercise2() {
	fmt.Println("**EXERCISE 2**")
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := toJSON(p1)
	if err != nil {
		log.Fatal("Error converting JSON: ", err)
	}

	fmt.Println(string(bs))
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
