package main

import (
	"fmt"
)

type Person struct {
	First string
	Last  string
	Age   int
}

func (p Person) FullName() string {
	return p.First + " " + p.Last
}

func (p Person) Greeting() string {
    return "My name is " + p.Last + ". " + p.First + " " + p.Last + "."
}

type DoubleZero struct {
	Person // embedded type
    First string
	LicenceToKill bool
}

func (dz DoubleZero) Greeting() string {
    return "I am agent 007."
}


func main() {

	p1 := DoubleZero{
		Person: Person{
			First: "James",
			Last:  "Bond",
			Age:   20,
		},
        First: "Double Zero Seven",
		LicenceToKill: true,
	}

    // fmt.Println(p1.Person.FullName())
	fmt.Println(p1.FullName()) // James Bond
    fmt.Println(p1) // {{James Bond 20} Double Zero Seven true}
    fmt.Println(p1.First, p1.Person.First) // Double Zero Seven James
    
    fmt.Println("-------------------")
    
    fmt.Println(p1.Greeting()) // I am agent 007.
    fmt.Println(p1.Person.Greeting()) // My name is Bond. James Bond.
    
    // outer type (DoubleZero) overrides inner type (Person)
}
