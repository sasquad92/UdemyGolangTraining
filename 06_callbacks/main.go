package main

import "fmt"

// here "callback" is just a name of func
func visit(numbers []int, callback func(int)) {
    for _, n := range numbers {
        callback(n)
    }
}


func filter(numbers []int, callback func(int) bool) []int {
    xs := []int{} // or var cs []int
    for _, n := range numbers {
        if callback(n) {
            xs = append(xs, n)
        }
    }
    return xs
}



func main() {
    // callback - passing func as an argument
    visit([]int{1, 2, 3, 4}, func(n int) {
        fmt.Println(n)
    })
    
    
    xs := filter([]int{1, 2, 3, 4}, func(n int) bool {
        return n > 1
    })
    
    fmt.Println(xs)
}