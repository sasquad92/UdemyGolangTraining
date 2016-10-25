package main

import (
	"fmt"
)

/*
CHALLANGE
Execute 100 factorial comp. concurrently and in parrarel.
Use the "pipeline" pattern to accomplish this.
*/

func main() {

	for n := range fact(generate()) {
		fmt.Println(n)
	}
}

func factorial(x int) int {
	sum := 1
	for i := 1; i <= x; i++ {
		sum *= i
	}
	return sum
}

func fact(in chan int) chan int {
	out := make(chan int)

	go func() {
		for n := range in {
			out <- factorial(n)
		}
		close(out)
	}()
	return out
}

func generate() chan int {
	out := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				out <- j
			}
		}
		close(out)
	}()
	return out
}
