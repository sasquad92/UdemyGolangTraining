package main

import (
    "fmt"
    "github.com/sasquad92/udemyTraining/02_packages/rectanglePkg"
)

func main() {
  
  //r := rectanglePkg.NewRect(3, 4) // error
  r := rectanglePkg.NewRectangle(3, 4) // ok
  
  fmt.Println("rectangle: ", r)
  fmt.Println("rectangle area: ", r.Area())
  fmt.Println("rectangle circ: ", r.Circ())
}
