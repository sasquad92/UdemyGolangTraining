package main

import (
	"fmt"
)

func main() {
	/*
		// Set up the pipeline
		c := gen(2, 3)
		out := sq(c)

		// Consume the output
		fmt.Println(<-out) // 4
		fmt.Println(<-out) // 9
	*/

	for n := range sq(gen(2, 3)) {
		fmt.Println(n)
	}
}

//    <- == receive only
func gen(nums ...int) chan int { // func gen(nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()

	return out
}

func sq(in chan int) chan int { // func sq(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}
