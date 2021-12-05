package main

import (
	"fmt"
	"os"
)

type Reader interface {
	Read(b []byte) (n int, err error)
}

type Writer interface {
	Write(b []byte) (n int, err error)
}

type ReadWriter interface {
	Reader
	Writer
}

func main() {
	// os.Stdout implements Writer
	var w Writer = os.Stdout

	i, err := fmt.Fprintf(w, "Hello, writer!!!\n")
	if nil != err {
		fmt.Println(err)
	}
	fmt.Println(i)
}
