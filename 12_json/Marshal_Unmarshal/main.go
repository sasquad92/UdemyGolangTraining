package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First       string `json:"first name"`
	Last        string
	Age         int
	notExported int    // low case -> not exported; we can also use `json:"-"`
}

func main() {

	p1 := Person{"James", "Bond", 20, 007}

	bs, _ := json.Marshal(p1)

	fmt.Println(bs)
	fmt.Printf("%T \n", bs)
	fmt.Println(string(bs))
    
    var p2 Person
    fmt.Println(p2)
    
    json.Unmarshal(bs, &p2)
    fmt.Println(p2)

}
