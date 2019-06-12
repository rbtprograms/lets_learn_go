package main

import "fmt"

func main() {
	// exercise1()
	exercise2()
	exercise3()
	// exercise4()
	// exercise5()
	// exercise6()
}

func exercise1() {
	for i := 1; i <= 10000; i++ {
		fmt.Println(i)
	}
}

func exercise2() {
	for i := 65; i <= 90; i++ {
		for j := 0; j <= 2; j++ {
			fmt.Printf("%v\t%#U\n", i, i)
		}
	}
}

func exercise3() {
	bd := 1990
	for bd < 2020 {
		fmt.Println(bd)
		bd++
	}
}

func exercise4() {
	x := 100
	fmt.Printf("%d\t%b\t%#x\n", x, x, x)
	y := x << 1
	fmt.Printf("%d\t%b\t%#x\n", y, y, y)

}

var y string

func exercise5() {
	y = `Totoro is
	the best ever`
	fmt.Print(y)
}

const (
	a = 2016 + iota
	b = 2016 + iota
	c = 2016 + iota
	d = 2016 + iota
)

func exercise6() {
	fmt.Printf("%v\t%v\t%v\t%v\t", a, b, c, d)
}
