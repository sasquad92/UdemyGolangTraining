package main

import (
    "fmt"
)

func makeGreeter() func() string {
    return func() string {
        return "Hello World!"
    }
}


func main() {
    
    greeting := func() {
        fmt.Println("func expression - asigning func to variable")
    }
    
    greeting()
    
    greet := makeGreeter()
    
    // type of greet is func() string
    fmt.Printf("%T \n", greet)
    // type of greet() is string
    fmt.Printf("%T \n", greet())
    fmt.Println(greet())
}