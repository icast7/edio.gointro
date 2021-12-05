package main

import (
	"fmt"
)

func GetSum(array []int) int {
	sum := 0
	for i := 0; i < len(array); i++ {
		sum = sum + array[i]
	}
	return sum
}

func main() {
	values := []int{
		9, 4, 5, 10, 3, 11, 99, 2,
	}
	result := GetSum(values)
	fmt.Printf("Result: %d", result)
}
