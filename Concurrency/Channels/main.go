package main

import "fmt"

func worker(done chan bool) {
	fmt.Println("Working..")
	for i := 0; i < 2; i++ {
		fmt.Println("Super", i)

	}
	done <- false
}

func main() {
	done := make(chan bool)
	go worker(done)

	fmt.Println("Goroutine finished")
}

// done <- true signals that the worker() is done.
// <-done waits until a value is received, preventing the main function from exiting early.
// Using false instead of true works too—it's just a signal. But true is more intuitive for "done."
