package main

import (
	"fmt"
)

type User struct {
	Name string
}

func main() {
	u := User{Name: "Leto"}

	fmt.Println(u.Name)
	Modify(&u)
	fmt.Println(u.Name)
}

func Modify(u *User) {
	*u = User{Name: "Paul"}
}
