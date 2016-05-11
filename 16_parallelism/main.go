package main

import (
	"fmt"
	"runtime"
	"sync"
)

// init runs before code starts executing
// and does some setup
func init() {
	// using all the CPUs
	// no longer need to include after go 1.5 (by defoult using more than 1 core)
	runtime.GOMAXPROCS(runtime.NumCPU())
}

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}

func foo() {

	for i := 0; i < 20; i++ {
		fmt.Println("foo", i)
	}
	wg.Done()
}

func bar() {

	for i := 0; i < 20; i++ {
		fmt.Println("bar", i)
	}
	wg.Done()
}
