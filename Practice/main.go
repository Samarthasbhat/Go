package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		fmt.Println("Sending....")
		ch <- "Hello"

		fmt.Println("Sent!") // There is no sync for this, added  sleep at end of main func
	}()

	time.Sleep(3 * time.Second)
	fmt.Println("Receiving...")
	msg := <-ch
	fmt.Println("Received:", msg)
	time.Sleep(5 * time.Second)
}
