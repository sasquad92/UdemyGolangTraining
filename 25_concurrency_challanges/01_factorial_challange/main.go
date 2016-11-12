package main

import (
	"fmt"
	"sync"
)

/*
CHALLANGE #1
Change the code to execute 1000 factorial computations concurrently and in parallel.
Use the "fan in / fan out" pattern to accomplish this.

CHALLANGE #2
While running the factorial computations, try to find how much of your resources are being used.
Post the percentage of your resources being used to this discussion: https://goo.gl/BxKnOL
~6.5% cpu
*/

func main() {
	in := generate()

	// FAN OUT - multiple functions reading from one channel
	c0 := fact(in)
	c1 := fact(in)
	c2 := fact(in)
	c3 := fact(in)
	c4 := fact(in)
	c5 := fact(in)
	c6 := fact(in)
	c7 := fact(in)
	c8 := fact(in)
	c9 := fact(in)

	// FAN IN - multiple channels writing to the same channel
	var y int

	for n := range merge(c0, c1, c2, c3, c4, c5, c6, c7, c8, c9) {
		y++
		fmt.Println(y, "\t", n)
	}
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

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
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
		for i := 0; i < 100000; i++ {
			for j := 40; j < 50; j++ {
				out <- j
			}
		}
		close(out)
	}()
	return out
}
