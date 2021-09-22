package main

import "fmt"

func main() {
	ch := make(chan int)
	close(ch)
	for v := range ch {
		fmt.Println(v)
	}

	fmt.Printf("TODO... %+v", ch)
}
