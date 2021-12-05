package main

import (
	"fmt"
)

type User struct {
	FirstName, LastName string
}

// Pointer as receiver
func (u *User) Greeting() string {
	return fmt.Sprintf("Dear %s %s", u.FirstName, u.LastName)
}

func main() {
	u := &User{"John", "Doe"}
	fmt.Println(u.Greeting())
}
