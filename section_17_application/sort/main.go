package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{4, 7, 2, 53, 88, 1, 98, 6, 32}
	xs := []string{"Subby", "Bubby", "Totobutts", "Bubbles", "Reefer", "ButtButts"}

	fmt.Println(sort.IntsAreSorted(xi))
	sort.Ints(xi)
	fmt.Println(sort.IntsAreSorted(xi))

	fmt.Println(sort.StringsAreSorted(xs))
	sort.Strings(xs)
	fmt.Println(sort.StringsAreSorted(xs))
	 
	fmt.Println(xi)
	fmt.Println(xs)
}
