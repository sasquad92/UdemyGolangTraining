package main

import (
	"fmt"
)

func incrementatior() chan int { // we can use <-chan int - it specifies the channel dirrection; if no direction is given, the channel is bidirectional
	channel := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			channel <- i
		}
		close(channel)
	}()

	return channel
}

func puller(c chan int) chan int {
	outChannel := make(chan int)

	go func() {
		var sum int
		for n := range c {
			sum += n
		}
		outChannel <- sum
		close(outChannel)
	}()
	return outChannel
}

func main() {
	chan1 := incrementatior()
	chan2 := puller(chan1)

	for m := range chan2 {
		fmt.Println(m)
	}
}

//    sum
// 1 - 1
// 2 - 3
// 3 - 6
// 4 - 10
// 5 - 15
// 6 - 21
// 7 - 28
// 8 - 36
// 9 - 45
