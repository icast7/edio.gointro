package main

import (
	"fmt"
	"math"
)

// Using alias
type MyVeryOwnFloat float64

// Adding method to MyVeryOwnFloat type
func (f MyVeryOwnFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyVeryOwnFloat(-math.SqrtPi)
	fmt.Println(f.Abs())
}
