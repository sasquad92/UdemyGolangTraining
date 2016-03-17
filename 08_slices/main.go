package main

import (
    "fmt"
)

func main() {
    /*
    myArray := [10]int{} // array
    mySlice := []int{} // slice // without specyfing cap and len
    
    fmt.Printf("%T, %v \n", myArray, myArray) // zeros
    fmt.Printf("%T, %v \n", mySlice, mySlice) // empty
    
     x := make([]int, 2, 7) // type, length, capacity
     y := new([7]int)[0:2] // capacity, tpye, length
     
     z := make([]int, 2) // cap and len is 2
     
     fmt.Printf("%T, %v \n", x, x) // slices with two zeros
     fmt.Printf("%T, %v \n", y, y)
     fmt.Printf("%T, %v \n", z, z)
     
     for i := 0; i < 10; i++ {
         x = append(x, i) // adding items to slice
         fmt.Println("Length:", len(x), "Capacity:", cap(x), "Value:", x[i])
     }
    fmt.Println(x) // 12 items - 2 zeros was there before append
    
    z = append(z, x[6:]...)
    fmt.Println(z)
    z = append(x, z...)
    fmt.Println(z)
    */
    
    // slice of slices
    
    students := make([][]string, 0) // var students []string
    
    student1 := []string{"Foster", "Nathan", "100.00", "74.00"}
    student2 := []string{"Gomez", "Lisa", "92.00", "96.00"}
    student3 := []string{"McMillan", "Peter", "81.50", "56.00"}
    student4 := []string{"Jonson", "Ruphert", "93.30", "77.00"}
    
    students = append(students, student1)
    students = append(students, student2)
    students = append(students, student3)
    students = append(students, student4)
    
    fmt.Println(students)
}