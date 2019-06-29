package main

import (
	"fmt"
	"sort"
)

type byAge []person
func (a byAge) Len() int           { return len(a) }
func (a byAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byAge) Less(i, j int) bool { return a[i].age < a[j].age }

type byName []person
func (n byName) Len() int           { return len(n) }
func (n byName) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }
func (n byName) Less(i, j int) bool { return n[i].first < n[j].first }

type person struct {
	first string
	age   int
}

func main() {
	p1 := person{"Bobby", 29}
	p2 := person{"Sophia", 27}
	p3 := person{"Caity", 32}
	p4 := person{"Injoong", 26}
	people := []person{p1, p2, p3, p4}

	fmt.Println(people)
	sort.Sort(byAge(people))
	fmt.Println(people)
	sort.Sort(byName(people))
	fmt.Println(people)
}
