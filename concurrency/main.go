package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func main() {
	t := tree.New(1)
	ch := make(chan int)
	close(ch)
	for v := range ch {
		fmt.Println(v)
	}
	fmt.Printf("Tree: %+v", t)
	fmt.Printf("TODO... %+v", ch)
}
