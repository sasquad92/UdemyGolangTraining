package main

import (
	"fmt"
)

func main() {

	// 1. Create a program which prints “Hello World” to the terminal.
	//fmt.Println("Hello World")

	// 2. Modify the previous program so that instead of printing Hello World it prints Hello, my name is followed by your name.
	//fmt.Println("Hello, my name is Krzysiek")

	// 3. Create a program that prints to the terminal asking for a user to enter their name. Use fmt.Scan to read a user’s name
	// entered at the terminal. Print “Hello <NAME>” with <NAME> replaced with what the user entered at the terminal.
	/*var name string

	  fmt.Println("What's your name?")
	  fmt.Scan(&name)
	  fmt.Println("Hello ", name)*/

	// 4. Create a program that prints to the terminal asking for a user to enter a small number and a larger number. Print the
	// remainder of the larger number divided by the smaller number.
	/*var smallNumber int
	var largeNumber int

	fmt.Println("Enter a small number")
	fmt.Scan(&smallNumber)
	fmt.Println("Enter a big number")
	fmt.Scan(&largeNumber)

	fmt.Println("Remainder of ", smallNumber, " / ", largeNumber, " is ", largeNumber%smallNumber)*/

	// 5. Print all of the even numbers between 0 and 100.
	/*for i := 0; i <= 100; i++ {
	    if i%2 == 0 {
	        fmt.Println(i)
	    }
	}*/

	// 6. Write a program that prints the numbers from 1 to 100. But for multiples of three print "Fizz" instead of the number and
	// for the multiples of five print "Buzz". For numbers which are multiples of both three and five print "FizzBuzz".
	/*for i := 1; i <= 100; i++ {

	    if i%3 == 0 && i%5 == 0 {
	        fmt.Println("FizzBuzz")
	    } else if i%3 == 0 {
	        fmt.Println("Fizz")
	    } else if i%5 == 0 {
	        fmt.Println("Buzz")
	    } else {
	        fmt.Println(i)
	    }
	}*/

	// 7. If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples
	// is 23. Find the sum of all the multiples of 3 or 5 below 1000.
	var sum int
	for i := 0; i < 1000; i++ {
		if i%3 == 0 {
			sum += i
		} else if i%5 == 0 {
			sum += i
		}
	}

	fmt.Println("Sum is equal to ", sum) // 233168
}
