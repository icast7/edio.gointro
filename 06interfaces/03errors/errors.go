package main

import (
	"fmt"
	"time"
)

type MyVeryOwnError struct {
	When time.Time
	What string
}

func (e *MyVeryOwnError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

// Always return error
func run() error {
	return &MyVeryOwnError{
		time.Now(),
		"it failed miserably",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
