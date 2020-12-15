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
	if (len(a) > 0 && len(b) == 0) {
		return 0, errors.New("no B provided")
	}
	if (len(a) == 0 && len(b) > 0) {
		return 0, errors.New("no A provided")
	}
	distance := 0
	for i, letter := range a {
		if (string(letter) != string(b[i])) {
			distance++
		}
	}
	return distance, nil
}
