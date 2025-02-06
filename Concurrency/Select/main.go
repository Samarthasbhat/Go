package main

import (
	"fmt"
	// "time"
	"sync"
)

func main(){
	ch1 := make(chan string)
	ch2 := make(chan string)
	var wg sync.WaitGroup // Wait group to manage goroutines

	wg.Add(2) // We have 2 goroutines

	// Simulating concurrent operations
	go func(){
		// time.Sleep(2*time.Second)
		defer wg.Done()
		ch1 <- "Data from ch1"
	}()  // The () at the end immediately executes the function.

	go func(){
		// time.Sleep(1 * time.Second)
		defer wg.Done()
		ch2 <- "Data from ch2"
	}()

	//  Using select to receive from whiever channel is ready first

	for i:=0; i<2; i++{ // select listens only for ch1 or ch2
select {
case msg := <-ch1:
	fmt.Println("Received:", msg)
case msg := <-ch2:
	fmt.Println("Received:", msg)
     }
	}

	 wg.Wait()
}

