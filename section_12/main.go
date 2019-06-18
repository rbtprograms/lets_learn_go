package main

import "fmt"

func main() {
	exercise1()
	exercise2()
	exercise3()
	exercise4()
}

type person struct{
	first string
	last string
	faveFlavors []string
}

func exercise1() {
	p1 := person{
		first: "Bobby",
		last: "Bubbero",
		faveFlavors: []string{"Vanilla", "Rocky Road"},
	}
	p2 := person{
		first: "Sophia",
		last: "Subbero",
		faveFlavors: []string{"Chocolate"},
	}
	for i, v := range p1.faveFlavors {
		fmt.Println(i, v)
	}
	for i, v := range p2.faveFlavors {
		fmt.Println(i, v)
	}
}

func exercise2() {
	p1 := person{
		first: "Bobby",
		last: "Bubbero",
		faveFlavors: []string{"Vanilla", "Rocky Road"},
	}
	p2 := person{
		first: "Sophia",
		last: "Subbero",
		faveFlavors: []string{"Chocolate"},
	}

	m := map[string]person {
		p1.last:p1,
		p2.last:p2,
	}
	for _, v := range m {
		fmt.Println(v)
		fmt.Println(v.first)
		for i, val := range v.faveFlavors {
			fmt.Println(i, val)
		}
		fmt.Println("------")
	}
}

func exercise3() {
	type vehicle struct{
		doors int
		color string
	}
	type truck struct{
		vehicle
		fourWheel bool
	}
	type sedan struct{
		vehicle
		luxury bool
	}
	t := truck{
		vehicle: vehicle{
			doors: 4,
			color: "white",
		},
		fourWheel: true,
	}
	s := sedan{
		vehicle: vehicle{
			doors: 2,
			color: "silver",
		},
		luxury: false,
	}
	fmt.Println(t.doors)
	fmt.Println(s.vehicle.doors)
}

func exercise4() {
	b := struct{
		first string
		last string
		favRappers map[string]string
		favFoods []string
	}{
		first: "Bobby",
		last: "Thompson",
		favRappers: map[string]string{
			"Kanye": "Chicago",
			"Jay-Z": "Brooklyn",
		},
		favFoods: []string{"cheese", "sausage"},
	}
	fmt.Println(b.first)
	fmt.Println(b.favFoods)

	for k, v := range b.favRappers {
		fmt.Println(k, v)
	}

	for i, v := range b.favFoods {
		fmt.Println(i, v)
	}
}