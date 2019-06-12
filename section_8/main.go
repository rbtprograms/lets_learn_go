package main

import "fmt"

func main() {
	// exercise1()
	exercise2()
	exercise3()
	exercise4()
	exercise5()
	exercise6()
	exercise7()
	exercise8()
	exercise9()
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
	bd := 1990
	for {
		if bd == 2020 {
			break
		}
		fmt.Println(bd)
		bd++
	}
}

func exercise5() {
	for i := 10; i <= 100; i++ {
		fmt.Printf("%v\n", i%4)
	}
}

var temp = 5

func exercise6() {
	if temp < 6 {
		fmt.Print("worked")
	}
}
func exercise7() {
	if temp > 6 {
		fmt.Print("worked")
	} else {
		fmt.Print("didnt")
	}
}
func exercise8() {
	switch {
	case false:
		fmt.Println("should not print")
	case true:
		fmt.Println("Will print")
	}
}
func exercise9() {
	favSport := "surfing"
	switch favSport {
	case "skiing":
		fmt.Println("skiing")
	case "surfing":
		fmt.Println("Hell yea brother")
	}
}
