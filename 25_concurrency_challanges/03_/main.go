package main

import (
	"fmt"
	// "sync"
	// "sync/atomic"
)

// var count int64
// var wg sync.WaitGroup

func main() {
	// wg.Add(2)
	chan1 := incrementator("1")
	chan2 := incrementator("2")
	// wg.Wait()
	var sum int

	for n := range chan1 {
		sum += n
	}

	for m := range chan2 {
		sum += m
	}

	fmt.Println("Final counter:", sum)
}

func incrementator(s string) chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 20; i++ {
			//atomic.AddInt64(&count, 1)
			out <- 1
			fmt.Println("Process: "+s+" printing:", i)
		}
		close(out)
	}()
	// wg.Done()
	return out
}

/*
CHALLANGE #1
Take the code from above and  change it to use channels instead of wait groups and atomicity
*/
