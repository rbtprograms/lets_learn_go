package main

import (
	"strconv"
	"strings"
	"unicode"
	"fmt"
)

func main() {
	fmt.Println(Valid("059"))
}

func Valid(check string) bool {
	if len(check) < 2 {
		return false
	}
	check = strings.ReplaceAll(check, " ", "")
	if len(check) <= 1 {
		return false
	}
	// make a new slice with all ints
	digits := make([]int, len(check))
	for i, n := range check {
		if !unicode.IsDigit(n) {
			return false
		}
		digits[i], _ = strconv.Atoi(string(n))
	}

	// double every other digit in the new slice
	for i := len(digits) - 2; i >= 0; i -= 2 {
		digits[i] *= 2
		if digits[i] > 9 {
			digits[i] -= 9
		}
	}
	// sum all those digits
	var sum int
	for _, n := range digits {
		sum += n
	}
	return sum%10 == 0
}
