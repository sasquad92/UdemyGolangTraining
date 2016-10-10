package main

import (
	"fmt"
	//"time"
)

func main() {
	c := make(chan int) // unbuffered channel

	// unbuffered channels block
	// go func() {
	//     for i := 0; i < 10; i++ {
	//         c <-i
	//     }
	// }()

	// go func() {
	//     for {
	//         fmt.Println(<-c)
	//     }
	// }()

	// time.Sleep(time.Second)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c) // can't put values on to the channel
	}()

	// waiting to receive data | waiting with main execution
	for n := range c {
		fmt.Println(n)
	}
}
