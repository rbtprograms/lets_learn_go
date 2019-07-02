package main

import "fmt"

func main() {
	exercise1()
	exercise2()
	exercise3()
	exercise4()
	exercise5()
	exercise6()
	exercise7()
	exercise8()
	exercise9()
	exercise10()
}

func exercise1() {
	a := [5]int{1, 2, 3, 4, 5}
	for i, v := range a {
		fmt.Println(i, v)
	}
	fmt.Printf("%T ", a)
}

func exercise2() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i, v := range s {
		fmt.Println(i, v)
	}
	fmt.Printf("%T ", s)
}

func exercise3() {
	s := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(s[2:5])
	fmt.Println(s[3:6])
	fmt.Println(s[4:7])
	fmt.Println(s[5:8])
	fmt.Println(s[6:9])
}

func exercise4() {
	s := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	s = append(s, 52)
	fmt.Println(s)
	s = append(s, 53, 54, 55)
	fmt.Println(s)
	s = append(s, []int{56, 57, 58, 59}...)
	fmt.Println(s)
}

func exercise5() {
	s := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	s = append(s[:3], s[6:]...)
	fmt.Println(s)
}

func exercise6() {
	s := make([]string, 50, 50)
	s = []string{}
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))
}
func exercise7() {
	s := [][]string{[]string{"Bobby", "Sophia", "yith"}, []string{"Cait", "Mom", "Dad"}}
	fmt.Println(s)

	for i, v := range s {
		fmt.Println("item: ", i)
		for j, v := range v {
			fmt.Printf("\tindex: %v\tvalue: %v", j, v)
		}
	}
}
func exercise8() {
	m := map[string][]string{
		"Bobby":   []string{"Sophia", "BBQ", "hikes"},
		"Sophia":  []string{"Resting alut", "being a been", "being the best"},
		"Injoong": []string{"Hades", "Ruth", "Reggie"},
	}
	for k, v := range m {
		fmt.Printf("\n\tthe record for key: %v\n", k)
		for i, v := range v {
			fmt.Println("\t", i, v)
		}
	}
}
func exercise9() {
	m := map[string][]string{
		"Bobby":   []string{"Sophia", "BBQ", "hikes"},
		"Sophia":  []string{"Resting alut", "being a been", "being the best"},
		"Injoong": []string{"Hades", "Ruth", "Reggie"},
	}
	m["Arthur"] = []string{"Board Games", "Friends", "climbing"}
	for k, v := range m {
		fmt.Printf("\n\tthe record for key: %v\n", k)
		for i, v := range v {
			fmt.Println("\t", i, v)
		}
	}
}
func exercise10() {
	m := map[string][]string{
		"Bobby":   []string{"Sophia", "BBQ", "hikes"},
		"Sophia":  []string{"Resting alut", "being a been", "being the best"},
		"Injoong": []string{"Hades", "Ruth", "Reggie"},
	}
	delete(m, "Bobby")
	for k, v := range m {
		fmt.Printf("\n\tthe record for key: %v\n", k)
		for i, v := range v {
			fmt.Println("\t", i, v)
		}
	}
}
