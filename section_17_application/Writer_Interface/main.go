package main

import (
	"fmt"
	"os"
	"io"
)

//All of these implement the writer interface, so they
//can all achieve similar things.
func main() {
	fmt.Println("Hello world");
	fmt.Fprintln(os.Stdout, "Hello world");
	io.WriteString(os.Stdout, "Hello world");
}
