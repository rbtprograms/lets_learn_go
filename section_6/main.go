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
	x := 42
	fmt.Printf("%d\t%b\t%#x", x, x, x)
}

func exercise2() {
	a := 42 == 42
	b := 42 <= 43
	c := 42 >= 43
	d := 42 != 42
	e := 42 < 43
	f := 42 > 43
	fmt.Println(a, b, c, d, e, f)

}

func exercise3() {
	const a string = "hi"
	const b = 100
	fmt.Println(a, b)
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
