package main

import "fmt"

func main() {
	exercise1()
	// exercise2()
	// exercise3()
	// exercise4()
}

func foo() int {
	return 5
}
func bar() (int, string) {
	return 55, "hello"
}
func exercise1() {
	r1 := foo()
	r2a, r2b := bar()
	fmt.Println("r1", r1)
	fmt.Println("r2a", r2a)
	fmt.Println("r2b", r2b)
}
func exercise2() {
		fmt.Println("------")
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