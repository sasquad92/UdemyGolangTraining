package main

import "fmt"

func main() {
	fmt.Println("Hello World!")

	// https://golang.org/pkg/fmt/
	fmt.Println(42)
	fmt.Printf("%d - decimal \n%b - binary\n%x - hex \n%o - oct\n", 33, 33, 33, 33)
	fmt.Printf("%#x - %#X \n", 11, 11)

	fmt.Println("UTF=8")

	for i := 60; i < 200; i++ {
		fmt.Printf("%d \t %q \n", i, i)
	}

}
