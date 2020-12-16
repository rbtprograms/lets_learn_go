package main

import (
	"strconv"
	"fmt"
)

func main() {
	res1 := Convert(35)
	fmt.Println(res1)
	res2 := Convert(7)
	fmt.Println(res2)
	res3 := Convert(101)
	fmt.Println(res3)
}

// Convert changes numebrs into raindrop sounds
func Convert(n int) (res string) {
	if n%3 == 0 {
		res += "Pling"
	}
	if n%5 == 0 {
		res += "Plang"
	}
	if n%7 == 0 {
		res += "Plong"
	}
	if res == "" {
		res += strconv.Itoa(n)
	}
	return
}
