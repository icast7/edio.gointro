package main

import "fmt"

func main() {
	var a [2]string // string array size 2
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	b := [2]string{"hello", "mundo!"}
	fmt.Printf("%q", b)

	c := [...]string{"hola", "world!"}
	fmt.Printf("%q", c)

	fmt.Printf("%s\n", a) // as string
	fmt.Printf("%q\n", a) // with quotes

	var x [2][3]string
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			x[i][j] = fmt.Sprintf("row %d - column %d", i+1, j+1)
		}
	}
	fmt.Printf("%q", x)

	var y [2]string
	// y[3] = "Will not compile"
	fmt.Println(y)
}
