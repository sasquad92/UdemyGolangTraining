package main

import (
	"fmt"
)

func main() {
	num := factorial(20)

	for n := range num {
		fmt.Println(n)
	}
}

func factorial(x int) chan int {
	sum := 1
	out := make(chan int)
	go func() {
		for i := 1; i <= x; i++ {
			sum *= i
		}
		out <- sum
		close(out)
	}()
	return out
}
