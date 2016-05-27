package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// to check data race you can run "go run -race main.go"

var wg sync.WaitGroup
var counter int
var mutex sync.Mutex


func main() {
	wg.Add(2)
	go incrementor("Foo:")
	go incrementor("Bar:")
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func incrementor(s string) {
	for i := 0; i < 20; i++ {
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		mutex.Lock()
        // x := counter
		// x++
        // counter = x
        counter++
		fmt.Println(s, i, "Counter: ", counter)
        mutex.Unlock()
	}
	wg.Done()
}
