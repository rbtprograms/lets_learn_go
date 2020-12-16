package main

import ("errors"; "fmt")

func main() {
	res, _ := Distance("ABCD", "ACBD")
	fmt.Println(res);
}

// Distance calculates the Hamming distance between two DNA strands
func Distance(a, b string) (int, error) {
	if (len(a) != len(b)) {
		return 0, errors.New("Different lengths")
	}
	distance := 0
	for i, letter := range a {
		if (letter != rune(b[i])) {
			distance++
		}
	}
	return distance, nil
}
