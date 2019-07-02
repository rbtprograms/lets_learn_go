package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{
		First: "Bobby",
		Last:  "Thompson",
		Age:   29,
	}
	p2 := person{
		First: "Sophia",
		Last:  "Fujiki",
		Age:   27,
	}
	people := []person{p1, p2}
	fmt.Println(people)
	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println("ERROR:", err)
	}
	fmt.Println(string(bs))

	//dont want to make a new file, this is the result of the first marshal
	jsonString := `[{"First":"Bobby","Last":"Thompson","Age":29},{"First":"Sophia","Last":"Fujiki","Age":27}]`
	bytes := []byte(jsonString)
	fmt.Printf("%T\t", jsonString)
	fmt.Printf("%T\t", bytes)

	var people2 []person
	err = json.Unmarshal(bytes, &people2)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("the data:", people)
	for i, v := range people {
		fmt.Println("\nINDEX:", i)
		fmt.Println("\nVALUES:", v.First, v.Last, v.Age)
	}
}
