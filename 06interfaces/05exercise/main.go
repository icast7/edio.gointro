package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %g", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	z := 1.0
	for i := 0; i < 10; i++ {
		z = z - ((z*z)-x)/(2*z)
	}
	return z, nil
}

func main() {
	result, err := Sqrt(4.0)
	fmt.Printf("Sqrt: %+v, Error: %+v\n", result, err)

	result, err = Sqrt(-4.0)
	fmt.Printf("Sqrt: %+v, Error: %+v\n", result, err)
}
