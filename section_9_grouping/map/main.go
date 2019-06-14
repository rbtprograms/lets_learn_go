package main

import (
	"fmt"
)

//maps are kinda like objects, fast lookup, no order,
//key value pairing
func main() {
	m := map[string]int{
		"Bobby":  29,
		"Sophia": 27,
	}
	fmt.Println(m)
	fmt.Println(m["Sophia"])

	//the comma ok idiom
	if v, ok := m["Bobby"]; ok {
		fmt.Println("THE IF PRINT: ", v)
	}

	if v, ok := m["Injoong"]; ok {
		fmt.Println("SECOND IF PRINT: ", v)
	}

	//add to the map
	m["Injoong"] = 25
	if v, ok := m["Injoong"]; ok {
		fmt.Println("THIRD IF PRINT: ", v)
	}

	for k, v := range m {
		fmt.Println(k, v)
	}

	xi := []int{4,5,6,7,8,9}
	for i, v := range xi {
		fmt.Println(i, v)
	}

	//delete from the map
	delete(m, "Bobby")
	fmt.Println(m)

	if v, ok := m["Sophia"]; ok {
		fmt.Println("value: ", v)
		delete(m, "Sophia")
	}

}
