package main

import (
	"fmt"
)

func main() {

	m := make(map[string]int)
    n := map[string]int{"foo": 1, "bar": 2}
    var o = make(map[string]int)
    //var o = make(map[string]int, 100) //<- with optional capacity
    //var o = map[string]int{} //<- ok
    //var o map[string]int <- error, nil map

    // adding key-value pairs
	m["v1"] = 23
	m["v2"] = 44

    n["v1"] = 23
	n["v2"] = 44
    
    o["v1"] = 23
	o["v2"] = 44
    
	fmt.Println(m)
    fmt.Println(n)
    fmt.Println(o)
    
    // value, is-there-key
	_, isThere := m["v3"]
	_, isThere2 := m["v2"]

	fmt.Println(isThere)
	fmt.Println(isThere2)
    
    fmt.Println("------")
    
    delete(m, "v2")
    
	v, isThere3 := m["v2"]

	fmt.Println(isThere3, v)
}
