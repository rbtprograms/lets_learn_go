package main

import "fmt"

func main() {
	fmt.Println(Difference(10))
}
// Difference finds the difference between square of sums and sum of squares
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}

// SumOfSquares squares all nums <= int, sums, and returns
func SumOfSquares(n int) (sum int) {
	for i := 1; i <= n; i++ {
		sum += i * i
	}
	return
}

// SquareOfSum sums an int and squares it
func SquareOfSum(n int) (sum int) {
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum * sum
}
