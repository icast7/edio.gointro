package main

import (
	"fmt"
	"strings"
)

// Using as alias
type MyVeryOwnString string

func (s MyVeryOwnString) Uppercase() string {
	return strings.ToUpper(string(s))
}

func main() {
	fmt.Println(MyVeryOwnString("testToUpper").Uppercase())
}
