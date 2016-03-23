package main

import (
	"encoding/json"
	"fmt"
    "strings"
    "os"
)

type Person struct {
	First       string `json:"first name"`
	Last        string
	Age         int
	notExported int    // low case -> not exported; we can also use `json:"-"`
}

func main() {

	p1 := Person{"James", "Bond", 20, 007}

	// json.NewEncoder(os.Stdout).Encode(p1)
    
    encoder := json.NewEncoder(os.Stdout)
    encoder.Encode(p1)
    
    
    var p2 Person
    reader := strings.NewReader(`{"first name":"James","Last":"Bond","Age":20}`)
    
    json.NewDecoder(reader).Decode(&p2)
    
    fmt.Println(p2)
}
