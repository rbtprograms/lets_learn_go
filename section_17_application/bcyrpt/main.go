package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := `password123`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bs)
	err = bcrypt.CompareHashAndPassword(bs, []byte("hi"))
	if err != nil {
		fmt.Println("Unable to log in")
		return
	}
	fmt.Println("You logged in")
}
