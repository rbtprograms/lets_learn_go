package main

import (
	"fmt"
	"log"
	"encoding/json"
	"errors"
)

func main() {
	exercise1()
	exercise2()
	exercise3()
	exercise4()
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

type customErr struct {
	err error
}
func (c customErr) Error() string {
	return fmt.Sprintf("ERROR DOING THING: %v %v", c.err)
}
func foo(e error) {
	fmt.Println(e)
}

func exercise3() {
	fmt.Println("**EXERCISE 3**")

	c1 := customErr{errors.New("Erorr")}
	foo(c1)
}

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		se := fmt.Errorf("Error making square root: %v", f)
		return 0, sqrtError{"50.2289 N", "99.4656 W", se}
	}
	return 42, nil
}

func exercise4() {
	fmt.Println("**EXERCISE 4**")

	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
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
