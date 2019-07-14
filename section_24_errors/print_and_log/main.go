package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	_, err := os.Open("non-existent-file.txt")
	if err != nil {
		//try all!
		fmt.Println("err happened", err)
		log.Println("err happened", err)
		//logs something and then calls os.Exit(1)
		// log.Fatal(err)
		//logs something and then calls panic
		// log.panic(err)
		// panic(err)
	}
	test()
}

func test() {
	f, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	log.SetOutput(f)

	f2, err := os.Open("non-existent-file.txt")
	if err != nil {
		log.Println("FAILURE:", err)
	}
	defer f2.Close()
	fmt.Println("Error occured, check log.txt")
}
