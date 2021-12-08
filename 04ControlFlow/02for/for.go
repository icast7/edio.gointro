package main

import "fmt"

func main() {
	sum0 := 0
	for i := 0; i < 10; i++ {
		sum0 += i
	}
	fmt.Println(sum0)

	// iterate as long as sum1 is less than 1000
	sum1 := 1
	for sum1 < 1000 {
		sum1 += sum1
	}
	fmt.Println(sum1)

	// alternative to while
	sum2 := 1
	for sum2 < 1000 {
		sum2 += sum2
	}
	fmt.Println(sum2)

	// infinite loop
	for {
		fmt.Println(".")
	}
}
