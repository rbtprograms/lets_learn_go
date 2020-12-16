package main

// Source: exercism/problem-specifications
// Commit: 0d882ed scrabble-score: Apply new "input" policy
// Problem Specifications Version: 1.1.0

type scrabbleTest struct {
	input    string
	expected int
	wordModifier int
	letterModifiers []int
}

var scrabbleScoreTests = []scrabbleTest{
	{"a", 1, 1, []int{1}},                           // lowercase letter
	{"A", 1, 1,[]int{1}},                           // uppercase letter
	{"f", 12, 1,[]int{3}},                           // valuable letter, triple letter score
	{"at", 2, 1,[]int{1,1}},                          // short word
	{"zoo", 64, 2,[]int{3,1,1}},                        // short, valuable word, double word score, tripple letter on z
	{"street", 6, 1,[]int{1,1,1,1,1,1}},                      // medium word
	{"quirky", 32, 1,[]int{2,1,1,1,1,1}},                     // medium, valuable word, double letter on q
	{"OxyphenButazone", 123, 3,[]int{1,1,1,1,1,1,1,1,1,1,1,1,1,1,1}},            // long, mixed-case word
	{"pinata", 8, 1,[]int{1,1,1,1,1,1}},                      // english-like word
	{"", 0, 1,[]int{}},                            // empty input
	{"abcdefghijklmnopqrstuvwxyz", 87, 1,[]int{1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1}}, // entire alphabet available
}
