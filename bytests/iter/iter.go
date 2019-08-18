package iter

import "strings"

//Repeat repeats a charcter 5 times and returns the string
func Repeat(char string, times int) string {
	var repeated strings.Builder
	for i := 0; i < times; i++ {
		repeated.WriteString(char)
	}
	return repeated.String()
}
