package main

import (
	"fmt"
	"sync"
)

/*
CHALLANGE #1:
This code throws an error: fatal error: all goroutines are asleep - deadlock!
Fix this code!
*/

func main() {

	in := gen()

	//FAN OUT
	// Multiple functions from same channel until channel is cloesd
	// Distribute work across multiple functions (ten gorutines) that all read from in.

	xc := fanOut(in, 10)

	// FAN IN
	// multiplex multiple channels onto a single channel
	// merge the channels from c0 through c9 into a single channel

	for n := range merge(xc...) {
		fmt.Println(n)
	}
}

func gen() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			for j := 3; j < 13; j++ {
				out <- j
			}
		}
		close(out)
	}()
	return out
}

func fanOut(in <-chan int, n int) []<-chan int {
	//xc := make([]<-chan int, n) // oryginal code
	var xc []<-chan int // needed to be zero

	for i := 0; i < n; i++ {
		xc = append(xc, factorial(in))
	}
	return xc
}

func factorial(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- fact(n)
		}
		close(out)
	}()
	return out
}

func fact(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goruotines are
	// done. This must start after the wg.Add call.

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
