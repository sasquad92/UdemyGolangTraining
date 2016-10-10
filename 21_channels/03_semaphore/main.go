package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	// this way is wrong, programm will be blocked
	// nothing is waiting for channel c
	// we have to create gorutine
	// <- done
	// <- done
	// close(c)

	go func() {
		<-done // waiting for true
		<-done // waiting for true
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}
