package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk tree and send values to channel
func walk(t *tree.Tree, ch chan int) {
	recursiveWalk(t, ch)
	// close channel to allow range to finish
	close(ch)
}

// Walkrecursively through the tree and push values to the channel
// at each recursion
func recursiveWalk(t *tree.Tree, ch chan int) {
	if t != nil {
		recursiveWalk(t.Left, ch)
		ch <- t.Value
		recursiveWalk(t.Right, ch)
	}
}

func same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go walk(t1, ch1)
	go walk(t2, ch2)
	for {
		x1, ok1 := <-ch1
		x2, ok2 := <-ch2
		switch {
		case ok1 != ok2:
			// not the same size
			return false
		case !ok1:
			// both channels are empty
			return true
		case x1 != x2:
			// elements are different
			return false
		default:
			// keep iterating
		}
	}
}

func main() {
	ch := make(chan int)
	fmt.Println("- - - - - - - - - - - - - - ")
	go walk(tree.New(3), ch)
	for v := range ch {
		fmt.Printf("%d ", v)
	}
	fmt.Print("\n")
	fmt.Println("- - - - - - - - - - - - - - ")
	fmt.Println(same(tree.New(1), tree.New(1)))
	fmt.Println("- - - - - - - - - - - - - - ")
	fmt.Println(same(tree.New(1), tree.New(2)))
}
