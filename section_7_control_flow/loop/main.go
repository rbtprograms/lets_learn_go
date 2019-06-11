package main

import "fmt"

//no while loops in go
func main() {
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	x := 1
	for x < 10 {
		fmt.Println(x)
		x++
	}

	for {
		x++
		if x > 20 {
			break
		}
		if x % 3 == 0 {
			continue
		} else {
			fmt.Println(x)
			x++
		}
	}

	for i := 33; i <= 122; i++ {
		fmt.Printf("%v\t%#U\n", i, i)
	}
}
