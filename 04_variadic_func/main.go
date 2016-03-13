package main

import (
    "fmt"
)

// variadic func
func average (sf ...float64) float64 {
    fmt.Println(sf)
    fmt.Printf("%T \n", sf)
    
    var total float64
    
    for _, v := range sf {
        total += v
    }
    
    return total / float64(len(sf))
}

func main() {
    n := average(43, 56, 65, 81 ,12, 98, 99)
    fmt.Println(n)
    
    data := []float64{43, 56, 65, 81 ,12, 98, 99}
    
    // variadic arguments
    n2 := average(data...)
    fmt.Println(n2)
}