package main

import (
	"fmt"
	"math"
)

type Squere struct {
	side float64
}

type Circle struct {
	radius float64
}

func (s Squere) area() float64 {
	return s.side * s.side
}

func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}


// Shape is an interface that has a area() method
type Shape interface {
	area() float64
}

func info(s Shape) {
	fmt.Println(s)
	fmt.Println(s.area())
}

func main() {
	s := Squere{5}
	c := Circle{3}
    
// Becouse Squere and Circle have are() method they can be used as Shape
	info(s)
	info(c)
}

// Anything that has the area() methon implements Shape interface
