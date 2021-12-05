package main

import "fmt"

type User struct {
	FirstName, LastName string
}

func (u *User) Name() string {
	return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
}

type Namer interface {
	// The Namer interface is defined by the Name method
	Name() string
}

// Takes the interace as a param
func Greet(n Namer) string {
	return fmt.Sprintf("Dear %s", n.Name())
}

func main() {
	u := &User{"John", "Doe"}
	fmt.Println(Greet(u))
}
