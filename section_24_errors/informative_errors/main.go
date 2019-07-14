package main

import (
	"fmt"
	"errors"
	"log"
)

type squareError struct {
	length int
	width int
	err error
}

func (s squareError) Error() string {
	return fmt.Sprintf("a square error occured: %v %v %v", s.length, s.width, s.err)
}

//ErrNegativeSqRt common error being used elsewhere
var ErrNegativeSqRt = errors.New("cannot square root negative number")

func main() {
	fmt.Printf("%T\n", ErrNegativeSqRt)
	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		//one way to return errors
		//return 0, ErrNegativeSqRt
		//Errorf is a convenience function for calling errors.New()
		//return 0 ,fmt.Errorf("cannot sqrt negative number: %v", f)
		//the below is weird but this is how you make custom errors
		se := fmt.Errorf("square error: cannot square root negative: %v", f)
		return 0, squareError{5, 6, se}
	}
	return 42, nil
}

