package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

func main() {
	exercise1()
	exercise2()
	exercise3()
	exercise4()
	exercise5()
}

type user struct {
	First string
	Age   int
}

func exercise1() {
	u1 := user{
		First: "James",
		Age:   32,
	}
	u2 := user{
		First: "Moneypenny",
		Age:   27,
	}
	u3 := user{
		First: "M",
		Age:   54,
	}
	users := []user{u1, u2, u3}
	fmt.Println(users)
	PeopleJSON, err := json.Marshal(users)
	if err != nil {
		fmt.Println("ERROR:", err)
	}
	fmt.Println(string(PeopleJSON))
}

type person struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

func exercise2() {
	s := `[
		{
			"First":"James",
			"Last":"Bond",
			"Age":32,
			"Sayings":[
				"Shaken, not stirred",
				"Youth is no guarantee of innovation",
				"In his majesty's royal service"]
		},
		{
			"First":"Miss",
			"Last":"Moneypenny",
			"Age":27,
			"Sayings":[
				"James, it is soo good to see you",
				"Would you like me to take care of that for you, James?",
				"I would really prefer to be a secret agent myself."]
		},
		{
			"First":"M",
			"Last":"Hmmmm",
			"Age":54,
			"Sayings":[
				"Oh, James. You didn't.",
				"Dear God, what has James done now?",
				"Can someone please tell me where James Bond is?"]
		}
	]`
	fmt.Println(s)
	var people []person
	err := json.Unmarshal([]byte(s), &people)
	if err != nil {
		fmt.Println("ERR:", err)
	}
	fmt.Println(people)
	for i, v := range people {
		fmt.Printf("\nINDEX: %v and VALUE: %v\n", i, v)
	}
}

func exercise3() {
	p1 := person{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	p2 := person{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	p3 := person{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	people := []person{p1, p2, p3}
	fmt.Println(people)
	err := json.NewEncoder(os.Stdout).Encode(people)
	if err != nil {
		fmt.Println("There was an error:", err)
	}
}

func exercise4() {
	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	fmt.Println(sort.IntsAreSorted(xi))
	sort.Ints(xi)
	fmt.Println(sort.IntsAreSorted(xi))

	fmt.Println(sort.StringsAreSorted(xs))
	sort.Strings(xs)
	fmt.Println(sort.StringsAreSorted(xs))
}

type byAge []person

func (a byAge) Len() int           { return len(a) }
func (a byAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type byLast []person

func (n byLast) Len() int           { return len(n) }
func (n byLast) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }
func (n byLast) Less(i, j int) bool { return n[i].Last < n[j].Last }

func exercise5() {
	p1 := person{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	p2 := person{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	p3 := person{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	people := []person{p1, p2, p3}
	for _, v := range people {
		sort.Strings(v.Sayings)
	}
	fmt.Println(people)
	sort.Sort(byAge(people))
	fmt.Println("SORTED BY AGE:", people)
	sort.Sort(byLast(people))
	fmt.Println("SORTED BY LAST:", people)
}
