package main

import (
	"fmt"
    "math"
    "strconv"
)

func main() {

	// 1. Write a function which takes an integer. The function will have two returns. The first return should be the argument divided by 2.
	// The second return should be a bool that let’s us know whether or not the argument was even. For example:
	// half(1) returns (0, false)
	// half(2) returns (1, true)

	/*fmt.Println(half(1))
	fmt.Println(half(2))*/
    
    
    
    
    
    // 2. Modify the previous program to use a func expression.
    
    /*half := func (arg int) (int, bool) {
        var div int
        var odd bool

        div = arg / 2

        if arg%2 == 0 {
            odd = true
        } else {
            odd = false
        }
        return div, odd
    }
    
    fmt.Println(half(1))
    fmt.Println(half(2))*/
    
    
    
    
    
    // 3. Write a function with one variadic parameter that finds the greatest number in a list of numbers.
    
    /*number := greatest(1, 2, 3, 4, 5, 6, 1, 2, 3, 4)
    fmt.Println(number)*/
    
    
    
    
    
    // 4. What's the value of this expression: (true && false) || (false && true) || !(false && false)?
    
    // false || false || true -> true
    
    
    
    
    
    // 5. Write a function, foo, which can be called in all of these ways:
    //   func main() {
    //        foo(1, 2)
    //        foo(1, 2, 3)
    //        aSlice := []int{1, 2, 3, 4}
    //        foo(aSlice...)
    //        foo()
    //   }
    
    
    
    
        
    // 6. Find a problem at projecteuler.net then create the solution. Add a comment 
    // beneath your solution that includes a description of the problem. Upload your 
    // solution to github. Tweet me a link to your solution. 
    
    // Desc:
    /*
    Sum of digits sequence
    Problem 551
    Let a0, a1, a2, ... be an integer sequence defined by:

    a0 = 1;
    for n ≥ 1, an is the sum of the digits of all preceding terms.
    The sequence starts with 1, 1, 2, 4, 8, 16, 23, 28, 38, 49, ...
    You are given a10^6 = 31054319.

    Find a10^15.
    */
    x := math.Pow10(6)
    //y := math.Pow10(15)
    fmt.Println("a^6 = ", finder(x))
   // fmt.Println("a10^15 = ", finder(4))
}

// func for ex 1
func half(arg int) (int, bool) {
	var div int
	var odd bool

	div = arg / 2

	if arg%2 == 0 {
		odd = true
	} else {
		odd = false
	}
	return div, odd
}


// func for ex 3
func greatest(numbers ...int) int {
    var greatest int
    
    for _, n := range numbers {
        if n > greatest {
            greatest = n
        }
    }
    
    return greatest
}


// func for ex 5
func foo(x ...int) {
    
    if x == nil {
        fmt.Println("X does not exist")
    } else {
        fmt.Println(x)
    }
}

// project euler
// problem 551
// float 64 becouse math.Pow needs float64
func finder(n float64) float64 {
    var answer float64
    
    if n < 2 {
       answer = 1
   } else {
       answer = finder(n-1) + countDigits(finder(n-1))
   }

    return answer
}

func countDigits(number float64) float64 {
    var result float64
    numberOfDigits := len(strconv.FormatFloat(number, 'f', 0, 64)) // need to know how many digits a number has got
    //fmt.Println(number, numberOfDigits) // for debug
    
    for i := 1; i <= numberOfDigits; i++ {
        result += math.Floor(math.Mod(number, math.Pow10(i)) / math.Pow10(i-1)) // need to Floor becouse I don't want a resoult with ,
        //fmt.Println(result) // for debug
    }
    return result
}