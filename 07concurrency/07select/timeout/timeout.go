package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	response := make(chan *http.Response, 1)
	errors := make(chan *error)

	go func() {
		resp, err := http.Get("https://www.google.com/")
		if err != nil {
			errors <- &err
		}
		response <- resp
	}()
	for {
		fmt.Println("- - - - -")
		select {
		case r := <-response:
			fmt.Printf("%s\n", r.Body)
			// return
		case err := <-errors:
			log.Fatal(*err)
		case <-time.After(500 * time.Millisecond):
			fmt.Printf("Timed out!!!")
			return
		}
	}
}
