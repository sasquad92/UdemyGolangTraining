package main

import (
	"errors"
	"fmt"
	"log"
)

var ErrSqrtNegative = errors.New("sqrt of negative number")

func main() {
	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		// return 0, ErrSqrtNegative
		return 0, fmt.Errorf("sqrt of negative number %v", x) // printing a formatted error; this method is calling error.New
	}
	// implementation
	return 42, nil
}
