package main

import "fmt"

type User struct {
	FirstName, LastName string
}

func (u *User) Name() string {
	return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
}

type Customer struct {
	Id       int
	FullName string
}

// Name method for Customer interface
func (c *Customer) Name() string {
	return c.FullName
}

type Namer interface {
	Name() string
}

func Greet(n Namer) string {
	return fmt.Sprintf("Dearrr %s", n.Name())
}

func main() {
	u := &User{"John", "Doe"}
	fmt.Println(Greet(u))
	// Customer implements the Name() method required by the Namer interface
	c := &Customer{374230, "Cesc Fabregas"}
	fmt.Println(Greet(c))
}
