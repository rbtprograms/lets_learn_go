package main

import (
	"fmt"
)

//pointer denoted with &
//pointers pass ADDRESS IN MEMORY of value,
//not the value itself
//& gives address, * gives value of item at the address

func main() {
	x := 42
	y := "hi"
	fmt.Println(x)
	fmt.Println(&x)
	fmt.Println(y)
	fmt.Println(&y)

	//these types are different, value type and pointer to type
	fmt.Printf("\n%T\n", x)
	fmt.Printf("\n%T\n", &x)
	fmt.Printf("\n%v\n", *&x)
	fmt.Printf("\n%T\n", y)
	fmt.Printf("\n%T\n", &y)
	fmt.Printf("\n%v\n", *&y)
	
	z := &x
	fmt.Printf("\n%T\n", *z)
	fmt.Printf("\n%v\n", *z)

	*z = 43
	fmt.Println(x)
}
